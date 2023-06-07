package controllers

import (
	"io/ioutil"
	"net/http"

	"github.com/ShivamIITK21/cflockout-backend/db"
	"github.com/ShivamIITK21/cflockout-backend/helpers"
	"github.com/ShivamIITK21/cflockout-backend/models"
	"github.com/gin-gonic/gin"
)

func RefreshController() gin.HandlerFunc{
	
	return func(c* gin.Context) {

		client := &http.Client{}

		url := "https://codeforces.com/api/problemset.problems"
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error":"Error while generating request"})
		}

		res, err := client.Do(req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error":"Error while fetching problems from CF api"})
		}

		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error":"Error while reading problem body"})
		}

		problems, err := helpers.ProblemParser(body)
		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error":"Error while parsing the problems in the server"})
		}

		result := db.DB.Where("1 = 1").Delete(&models.Problem{})
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error":"could not empty database"})
		}
		
		result = db.DB.Create(&problems)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error":"could not fill database"})
		}

		c.JSON(200, gin.H{"chill" : "hai"})
	}
}