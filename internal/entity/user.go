package entity

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name             string `gorm:"column:name;not null"`
	Email            string `gorm:"column:email;not null"`
	Password         string `gorm:"column:password;not null"`
	Photo            string `gorm:"column:photo;not null"`
	VerificationCode string `gorm:"column:verification_code;not null"`
	IsVerified       bool   `gorm:"column:is_verified;not null"`

	Products []Product
}

func (impl *User) TableName() string {
	return "users"
}
