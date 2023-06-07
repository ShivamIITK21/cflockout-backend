package models

import (
	"gorm.io/gorm"
)

type Problem struct {
	ContestID uint     `json:"contestId" gorm:"PRIMARY_KEY;autoIncrement:false;NOT NULL"`
	Index     *string  `json:"index" gorm:"PRIMARY_KEY;autoIncrement:false;NOT NULL"`
	Name      *string  `json:"name" gorm:"NOT NULL"`
	Rating    uint     `json:"rating"`
	Tags      *[]string `json:"tags" gorm:"type:text[]"`
}

func MigrateProblem(db *gorm.DB) error {
	err := db.AutoMigrate(&Problem{})
	return err
}
