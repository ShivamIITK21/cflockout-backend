package models

type Submission struct {
	ContestId   uint    `json:"contestId"`
	Index       *string `json:"index"`
	TimeCreated int64   `json:"creationTimeSeconds"`
	Verdict     *string `json:"verdict"`
	Rating		uint	`json:"rating"`
	Name      	*string   `json:"name"`
	Author		*string  `json:"author"`
}
