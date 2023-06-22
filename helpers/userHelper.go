package helpers

import (
	"encoding/json"
	"errors"

	"github.com/ShivamIITK21/cflockout-backend/models"
)

//gets problem info from user.status
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

	for _, s := range submissionResult{

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
		
		name, ok:= prob["name"].(string)
		if(!ok) {
			return processedSubmissions, err
		}
		sub.Name = &name

		id, ok := prob["contestId"]
		if id != nil {
			sub.ContestId = uint(id.(float64))
		}
			
		rating := prob["rating"]
		if rating != nil{
			sub.Rating = uint(rating.(float64))
		} else {
			sub.Rating = uint(0)
		}
		
		processedSubmissions = append(processedSubmissions, sub)
	}

	return processedSubmissions, nil
}