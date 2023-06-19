package helpers

import (
	"net/http"
	"io/ioutil"
    "math/rand"
	"time"
	
	"github.com/ShivamIITK21/cflockout-backend/models"
	"github.com/ShivamIITK21/cflockout-backend/db"
)

func GetProblem() (models.Problem, error){
	var allProbs []models.Problem
	var prob models.Problem
	
	result := db.DB.Find(&allProbs)
	if result.Error != nil {
		return prob, result.Error
	}

	ind := rand.Intn(len(allProbs))

	return allProbs[ind], nil
}

func VerifyUser(problem models.Problem, user models.User) error{

	time.Sleep(80 * time.Second)
	var prob models.Problem
	client := &http.Client{}
	url := "https://codeforces.com/api/user.status?handle=" + *user.Username + "&count=10"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil{
		return err
	}
	submissions, err := ExtractSubmissionInfo(body)
	if err != nil{
		return err
	}

	compileError := false
	for _, s := range submissions{
		if(*s.Verdict=="COMPILATION_ERROR" && s.ContestId==prob.ContestID && s.Index==prob.Index){
			compileError = true
		}
	}

	if compileError{
		user_type := "user"
		user.UserType = &user_type
		token, _, err := GenerateTokens(*user.Username, *user.CFid, *user.UserType)
		if err != nil{
			return err
		}
		user.Token = &token
		db.DB.Create(&user)
	}

	return nil
}