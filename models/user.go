package models

type User struct {
	Username *string `json:"username" gorm:"PRIMARY_KEY;NOT NULL"`
	CFid     *string `json:"cfid" gorm:"NOT NULL"`
	Password *string `json:"password" gorm:"NOT NULL"`
	Token    *string `json:"token" gorm:"NOT NULL"`
	UserType *string `json:"userType" gorm:"NOT NULL"`
}
