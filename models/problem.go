package models

import (
	"gorm.io/gorm"
)

type Problem struct {
	ContestID uint     `json:"contestID" gorm:"PRIMARY_KEY;autoIncrement:false;NOT NULL"`
	Index     *string  `json:"index" gorm:"PRIMARY_KEY;autoIncrement:false;NOT NULL"`
	Name      *string  `json:"name" gorm:"NOT NULL"`
	Rating	  *string  `json:"rating"`
	Ptype     *string  `json:"type" gorm:"NOT NULL"`
	Tags      []string `json:"tags" gorm:"type:text"`
}

// func (m *Problem) PrimaryKey() []interface{} {
// 	return []interface{}{m.ContestID, m.Index}
// }

// func (m *Problem) BeforeCreate(db *gorm.DB) error {
// 	db.Statement.SetColumn("PrimaryKey", m.PrimaryKey())
// 	return nil
// }

func MigrateProblem(db *gorm.DB) error {
	err := db.AutoMigrate(&Problem{})
	return err
}
