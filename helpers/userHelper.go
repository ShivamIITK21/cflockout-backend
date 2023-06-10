package helpers

import(
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
		tempv := oneSub["verdict"].(string)
		sub.Verdict = &tempv
		sub.TimeCreated = int64(oneSub["creationTimeSeconds"].(float64))
		prob, ok := oneSub["problem"].(map[string]interface{})
		if !ok {
			return processedSubmissions, err
		}
		tempidx := prob["index"].(string)
		sub.Index = &tempidx
		sub.Rating = 0
		var name string
		sub.Name = &name
		sub.ContestId = uint(prob["contestId"].(float64))

		processedSubmissions = append(processedSubmissions, sub)
	}

	return processedSubmissions, nil
}