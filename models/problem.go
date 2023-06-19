package models

type Problem struct {
	ContestID uint    `json:"contestId" gorm:"PRIMARY_KEY;autoIncrement:false;NOT NULL"`
	Index     *string `json:"index" gorm:"PRIMARY_KEY;autoIncrement:false;NOT NULL"`
	Name      *string `json:"name" gorm:"NOT NULL"`
	Rating    uint    `json:"rating"`
	Tags      *string `json:"tags"`
}
