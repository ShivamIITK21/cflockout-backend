package controllers

import (
	"net/http"
	"strconv"

	"github.com/ShivamIITK21/cflockout-backend/db"
	"github.com/ShivamIITK21/cflockout-backend/helpers"
	"github.com/ShivamIITK21/cflockout-backend/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)


func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), err
}

func VerifyPassword(userPass string, providedPass string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(userPass), []byte(providedPass))
	var match bool = true
	var msg string = ""
	if err != nil {
		match = false
		msg = "Incorrect email or password"
	}
	return match, msg
}


func Login() gin.HandlerFunc{
	return func(c *gin.Context){
		var user models.User
		err := c.BindJSON(&user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error":"Could not bind json"})
			return 
		}
		
		var retUser models.User
		result := db.DB.First(&retUser, user.Username)
		
		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error":"user not in db"})
			return 
		}
		
		passValid, msg := VerifyPassword(*retUser.Password, *user.Password)
		if !passValid {
			c.JSON(http.StatusBadGateway, gin.H{"error":msg})
			return 
		}
		
		token, _, err := helpers.GenerateTokens(*retUser.Username, *retUser.CFid, *retUser.UserType)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error":"could not generate token"})
			return 
		}
		
		err = helpers.UpdateAllTokens(token, *retUser.Username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error":"could not update the token"})
			return 
		}
		
		c.JSON(http.StatusOK, gin.H{"token":token, "cfID": *retUser.CFid})
	}
} 
	
func Signup() gin.HandlerFunc{
	return func(c *gin.Context){
		var user models.User
		err := c.BindJSON(&user)
		if err!= nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error":"couldn't bind JSON object"})
			return
		}

		var dummy models.User
		result := db.DB.Where("username = ?", user.Username).First(&dummy)
		if result.Error == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error":"This username already exists"})
			return
		}
		
		pass, err := HashPassword(*user.Password)
		if err!= nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error":"couldn't hash password"})
			return
		}
		user.Password = &pass
		
		prob, err := helpers.GetProblem()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error":"error in generating random problen"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"problem": "codeforces.com/problemset/problem/" + strconv.FormatUint(uint64(prob.ContestID), 10) + "/" + *prob.Index})
			
		go helpers.VerifyUser(prob, user)
	}
}

func FindUser() gin.HandlerFunc{
	return func(c *gin.Context){
		username := c.Query("username")

		var finduser models.User
		db.DB.Where("username = ?", username).First(&finduser)
		if finduser.Username == nil {
			c.JSON(http.StatusOK, gin.H{"message":"Not found"})
			return 
		}
		c.JSON(http.StatusOK, gin.H{"message":"User " + username + " was found"})
	}
}


func VerifyToken() gin.HandlerFunc{
	return func(c *gin.Context){
		token := c.Request.Header.Get("token")

		if token == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error":"No token"})
			return
		}

		claims, msg := helpers.ValidateToken(token)
		if msg != "" {
			c.JSON(http.StatusBadRequest, gin.H{"error":msg})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message":msg, "cfid":claims.CFid, "username": claims.Username})
	}
}


func Logout() gin.HandlerFunc{
	return func(c *gin.Context){
		// var user models.User
		// err := c.BindJSON(&user)
		// if err != nil {
		// 	c.JSON()
		// }
	}
}