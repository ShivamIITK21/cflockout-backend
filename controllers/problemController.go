package controllers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/ShivamIITK21/cflockout-backend/db"
	"github.com/ShivamIITK21/cflockout-backend/helpers"
	"github.com/ShivamIITK21/cflockout-backend/models"
	"github.com/gin-gonic/gin"
)

func RefreshController() gin.HandlerFunc{
	
	return func(c* gin.Context) {

		user_type := c.GetString("user_type")
		if(user_type != "admin"){
			c.JSON(http.StatusBadRequest, gin.H{"error":"You don't have permissions"})
			return
		}

		client := &http.Client{}

		url := "https://codeforces.com/api/problemset.problems"
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error":"Error while generating request"})
			return
		}

		res, err := client.Do(req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error":"Error while fetching problems from CF api"})
			return 
		}

		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error":"Error while reading problem body"})
			return 
		}

		problems, err := helpers.ProblemParser(body)
		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error":"Error while parsing the problems in the server"})
			return 
		}

		result := db.DB.Where("1 = 1").Delete(&models.Problem{})
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error":"could not empty database"})
			return 
		}
		
		result = db.DB.Create(&problems)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error":"could not fill database"})
			return 
		}

		c.JSON(200, gin.H{"chill" : "hai"})
	}
}


func GetUserSolvedProblems() gin.HandlerFunc{

	return func(c* gin.Context){
		user := c.Query("user")

		client := &http.Client{}
		var url string
		if(user != ""){
			url = "https://codeforces.com/api/user.status?handle=" + user
		}else{
			url = "https://codeforces.com/api/user.status?handle=yuvrajKharayat"
		}

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error":"Error while generating request"})
			return 
		}

		res, err := client.Do(req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error":"Error while fetching status from CF api"})
			return 
		}

		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error":"Error while reading status body"})
			return 
		}

		submissions, err := helpers.ExtractSubmissionInfo(body)
		fmt.Println(*submissions[0].Author)
		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
			return
		}
		
		ACs := make(map[string]int)
		WAs := make(map[string]int)
		for _, s := range submissions{
			if *s.Verdict == "OK" {
				ACs[strconv.Itoa(int(s.ContestId)) + *s.Index] = 1
			} else {
				WAs[strconv.Itoa(int(s.ContestId)) + *s.Index] = 1
			}
		}

		var allProbs []models.Problem
		result := db.DB.Find(&allProbs)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error in retrieving from DB"})
			return
		}

		var allProbData []models.Submission
		for _, pData := range allProbs{
			var s models.Submission
			s.ContestId = pData.ContestID
			s.Index = pData.Index
			s.Rating = pData.Rating
			v := "NA"
			if ACs[strconv.Itoa(int(s.ContestId)) + *s.Index] == 1 && user != "" {
				v = "AC"
			} else if WAs[strconv.Itoa(int(s.ContestId)) + *s.Index] == 1 && user != ""{
				v = "WA"
			}
			s.Verdict = &v
			s.Name = pData.Name
			s.Author = &user
			allProbData = append(allProbData, s)
		}

		helpers.SortSubmissions(&allProbData)
		
		c.JSON(200, allProbData)
		
	}
}
