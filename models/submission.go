package models

type Submission struct {
	ContestId   uint    `json:"contestId"`
	Index       *string `json:"index"`
	TimeCreated int64   `json:"creationTimeSeconds"`
	Verdict     *string `json:"verdict"`
}
