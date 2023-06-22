package models

import (
	"gorm.io/datatypes"
)

type ProblemInfo struct {
	Task          Problem `json:"task"`
	Score         *string `json:"score"`
	FirstSolvedBy *string `json:"firstsolvedby"`
}

type SessionData struct {
	Participants *map[string]string                `json:"participants"`
	Problems     datatypes.JSONType[[]ProblemInfo] `json:"problems"`
}

type Lockout struct {
	SessionId   *string                         `json:"sessionid" gorm:"PRIMARY_KEY; NOT NULL"`
	SessionData datatypes.JSONType[SessionData] `json:"session_data"`
	StartTime	int64							`json:"start_time"`
	Duration	int64							`json:"duration"`
}
