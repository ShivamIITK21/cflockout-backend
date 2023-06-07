package helpers

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/ShivamIITK21/cflockout-backend/models"
)

func ProblemParser(rawData []byte) ([]models.Problem, error) {

	var responseData map[string]interface{}
	var processedProblems []models.Problem
	
	err := json.Unmarshal(rawData, &responseData)
	if err != nil {
		return processedProblems, err
	}
	responseData, ok := responseData["result"].(map[string]interface{})
	err = errors.New("error during typecast")
	if !ok {
		return processedProblems, err
	}
	
	allProblems, ok := responseData["problems"].([]interface{})
	if !ok {
		return processedProblems, err
	}

	for _, p := range allProblems{
		probModel, ok := p.(models.Problem)
		if !ok {
			fmt.Println("Sorry could not parse")
			continue
		}
		processedProblems = append(processedProblems, probModel)
	}

	return processedProblems, nil
}