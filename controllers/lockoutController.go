package controllers

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"

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
			c.JSON(http.StatusBadRequest, gin.H{"error" : "one or more participaants are needed"})
			return
		}

		if len(req.Ratings) != len(req.Score) {
			c.JSON(http.StatusBadRequest, gin.H{"error" : "Length of ratings and score not equal"})
			return
		}

		if len(req.Ratings) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error" : "Len of ratings should not be zero"})
			return
		}

		for _, username := range req.Participants {
			var usr models.User
			result := db.DB.Where("username = ?", username).First(&usr)
			if result.Error != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error" : "user not in db"})
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
			if(err != nil){
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
		db.DB.Create(&session)

		c.JSON(http.StatusOK, gin.H{"chill":"hai"})
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
			c.JSON(http.StatusBadRequest, gin.H{"error" : "Session id invalid"})
			return 
		}

		if status != "admin" {
			user_map := *lockout.SessionData.Data().Participants
			found := false
			for reg_user := range user_map{
				if user == reg_user {
					found = true
				}
			}
			if !found {
				c.JSON(http.StatusBadRequest, gin.H{"error":"You are not a participant of this lockout"})
				return 
			}
		}

		c.JSON(http.StatusOK, lockout)

	}
}
