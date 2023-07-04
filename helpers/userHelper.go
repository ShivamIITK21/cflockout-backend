package helpers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/ShivamIITK21/cflockout-backend/models"
)

// gets problem info from user.status
func ExtractSubmissionInfo(rawData []byte) ([]models.Submission, error) {

	var processedSubmissions []models.Submission
	var interfaceData map[string]interface{}

	err := json.Unmarshal(rawData, &interfaceData)
	if err != nil {
		return processedSubmissions, err
	}

	submissionResult, ok := interfaceData["result"].([]interface{})
	err = errors.New("error during typecast")
	if !ok {
		return processedSubmissions, err
	}

	for _, s := range submissionResult {

		oneSub, ok := s.(map[string]interface{})
		if !ok {
			return processedSubmissions, err
		}

		var sub models.Submission

		tempv, ok := oneSub["verdict"].(string)
		if !ok {
			return processedSubmissions, err
		}
		sub.Verdict = &tempv

		time, ok := oneSub["creationTimeSeconds"].(float64)
		if !ok {
			return processedSubmissions, err
		}
		sub.TimeCreated = int64(time)

		prob, ok := oneSub["problem"].(map[string]interface{})
		if !ok {
			return processedSubmissions, err
		}

		tempidx, ok := prob["index"].(string)
		if !ok {
			return processedSubmissions, err
		}
		sub.Index = &tempidx

		name, ok := prob["name"].(string)
		if !ok {
			return processedSubmissions, err
		}
		sub.Name = &name

		id := prob["contestId"]
		if id != nil {
			sub.ContestId = uint(id.(float64))
		}

		rating := prob["rating"]
		if rating != nil {
			sub.Rating = uint(rating.(float64))
		} else {
			sub.Rating = uint(0)
		}

		author, ok := oneSub["author"].(map[string]interface{})
		if !ok {
			return processedSubmissions, err
		}

		member, ok := author["members"].([]interface{})
		if !ok {
			return processedSubmissions, err
		}

		mem1, ok := member[0].(map[string]interface{})
		if !ok {
			return processedSubmissions, err
		}

		handle := mem1["handle"].(string)
		sub.Author = &handle
		processedSubmissions = append(processedSubmissions, sub)
	}

	return processedSubmissions, nil
}

func RequestSubmissions(user string, numberOfSubmissions int) ([]models.Submission, error) {

	client := &http.Client{}
	var url string
	if user != "" {
		url = "https://codeforces.com/api/user.status?handle=" + user + "&count=" + strconv.Itoa(numberOfSubmissions)
	} else {
		url = "https://codeforces.com/api/user.status?handle=yuvrajKharayat"
	}

	var submissions []models.Submission
	newErr := errors.New("error during extracting latest submissions")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return submissions, newErr
	}

	res, err := client.Do(req)
	if err != nil {
		return submissions, newErr
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return submissions, newErr
	}

	submissions, err = ExtractSubmissionInfo(body)
	if err != nil {
		return submissions, newErr
	}

	var ACsubmissions []models.Submission

	for _, submission := range submissions {
		if *submission.Verdict == "OK" {
			ACsubmissions = append(ACsubmissions, submission)
		}
	}

	return ACsubmissions, nil
}
