package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"uniqueIndex"`
	Password string `gorm:"type:varchar(255) not null"`
}

func (User) TableName() string {
	return "user"
}
 