package model

import (
	"context"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `gorm:"uniqueIndex;type:varchar(255) not null"`
	Password string `gorm:"type:varchar(255) not null"`
}

func (User) TableName() string {
	return "user"
}

type UserQuery struct {
	db *gorm.DB
}

func NewUserQuery(db *gorm.DB) *UserQuery {
	return &UserQuery{
		db: db,
	}
}

func (uq *UserQuery) Create(ctx context.Context, user *User) error {
	return uq.db.Create(user).Error
}

func (uq *UserQuery) FindByEmail(ctx context.Context, email string) (*User, error) {
	var user User
	err := uq.db.Where("email = ?", email).Find(&user).Error
	return &user, err
}
