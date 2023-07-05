package controllers

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/ShivamIITK21/cflockout-backend/db"
	"github.com/ShivamIITK21/cflockout-backend/helpers"
	"github.com/ShivamIITK21/cflockout-backend/models"
	"github.com/gin-gonic/gin"
	"gorm.io/datatypes"
)

type CreateDetails struct {
	Participants []string `json:"participants"`
	Ratings      []string `json:"ratings"`
	Score        []string `json:"score"`
	StartsIn   int64    `json:"start_time"`
	Duration     int64    `json:"duration"`
}

func GetRandomID() string {
	randomBytes := make([]byte, 15)
	rand.Read(randomBytes)
	return base64.URLEncoding.EncodeToString(randomBytes)[:15]
}

func CreateLockoutController() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req CreateDetails
		err := c.BindJSON(&req)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
			return
		}

		if len(req.Participants) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "one or more participaants are needed"})
			return
		}

		if len(req.Ratings) != len(req.Score) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Length of ratings and score not equal"})
			return
		}

		if len(req.Ratings) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Len of ratings should not be zero"})
			return
		}
		
		if req.Duration > 86400 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Lockout duration exceeds maximum limit of 1 day"})
			return
		}

		if req.Duration < 60 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Lockout duration can't be less than a minue"})
			return
		}

		if req.StartsIn < 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Lockout start time is in past"})
			return
		}

		for _, username := range req.Participants {
			var usr models.User
			result := db.DB.Where("c_fid = ?", username).First(&usr)
			if result.Error != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "user not in db"})
				return
			}
		}

		id := GetRandomID()
		var participants_score = make(map[string]string)
		for _, username := range req.Participants {
			participants_score[username] = "0"
		}

		var probInfo []models.ProblemInfo
		for i, rating := range req.Ratings {
			prob, err := helpers.GetProblemByRating(rating)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Error while generating problems"})
				return
			}
			empty := ""
			var current models.ProblemInfo
			current.Task = prob
			current.Score = &req.Score[i]
			current.FirstSolvedBy = &empty
			probInfo = append(probInfo, current)
		}

		var sessionData models.SessionData
		sessionData.Participants = &participants_score
		sessionData.Problems = datatypes.NewJSONType(probInfo)
		var session models.Lockout
		session.SessionId = &id
		session.SessionData = datatypes.NewJSONType(sessionData)
		fmt.Println(req.StartsIn, time.Now().Unix())
		session.StartsIn = req.StartsIn
		session.Duration = req.Duration
		session.Processing = true
		db.DB.Create(&session)

		go SessionHandler(*session.SessionId)		

		c.JSON(http.StatusOK, gin.H{"session_id": *session.SessionId})
	}
}


func SessionHandler(session_id string){
	var lockout models.Lockout
	db.DB.Where("session_id = ?", session_id).First(&lockout)
	fmt.Println("Waiting for lockout to start")
	for time.Now().Unix() < lockout.StartsIn {
		time.Sleep(time.Second * 60)
	}
	fmt.Println("Lockout Started!")
	start_time := time.Now().Unix()
	for{
		if(start_time + lockout.Duration + 120 <= time.Now().Unix()){
			break;
		}
		
		db.DB.Where("session_id = ?", session_id).First(&lockout)
		
		var allSubmissions []models.Submission
		for participant := range *lockout.SessionData.Data().Participants{
			newSubmissions, err := helpers.RequestSubmissions(participant, 10)
			if err != nil {
				db.DB.Model(&models.Lockout{}).Where("session_id = ?", session_id).Update("processing", false)
				return
			}
			allSubmissions = append(allSubmissions, newSubmissions...)
		}

		helpers.SortSubmissionsByTime(&allSubmissions)

		for _, submission := range allSubmissions{
			for i, problem := range lockout.SessionData.Data().Problems.Data() {
				if (problem.Task.ContestID == submission.ContestId) && (*problem.Task.Index == *submission.Index) && (*problem.FirstSolvedBy=="") {
					lockout.SessionData.Data().Problems.Data()[i].FirstSolvedBy = submission.Author
					for user, score := range *lockout.SessionData.Data().Participants{
						if user == *submission.Author{
							userScore, _ := strconv.Atoi(score)
							probScore, _ := strconv.Atoi(*problem.Score)
							newScore := strconv.Itoa(userScore + probScore)
							(*lockout.SessionData.Data().Participants)[user] = newScore
						}
					}	
				}
			}
		}	
		
		db.DB.Model(&lockout).Updates(lockout)
		
		time.Sleep(5 * time.Second)
	}
	
}


func LockoutController() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.GetString("username")
		status := c.GetString("user_type")

		session_id := c.Query("session_id")
		var lockout models.Lockout
		result := db.DB.Where("session_id = ?", session_id).First(&lockout)
		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Session id invalid"})
			return
		}

		if status != "admin" {
			user_map := *lockout.SessionData.Data().Participants
			found := false
			for reg_user := range user_map {
				if user == reg_user {
					found = true
				}
			}
			if !found {
				c.JSON(http.StatusBadRequest, gin.H{"error": "You are not a participant of this lockout"})
				return
			}
		}

		c.JSON(http.StatusOK, lockout)

	}
}

func GetRating() gin.HandlerFunc {
	return func(c *gin.Context) {

		cfid := c.Query("cfid")

		client := &http.Client{}
		url := "https://codeforces.com/api/user.info?handles=" + cfid
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error":"Error while generating request"})
			return 
		}

		res, err := client.Do(req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error":"Error while fetching user info"})
			return 
		}

		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error":"Error while reading info"})
			return 
		}

		var interfaceData map[string]interface{}
		err = json.Unmarshal(body, &interfaceData)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error":"Error while unmarshaling"})
			return 
		}

		userInfo, ok := interfaceData["result"].([]interface{})
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error":"Error while typecasting"})
			return 
		}

		userData, ok := userInfo[0].(map[string]interface{})
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error":"Error while reading data"})
			return 
		}

		rating, ok := userData["rating"]
		c.JSON(200, gin.H{"rating": rating})

	}
}
