package entity

import "time"

type User struct {
	ID               int64     `gorm:"primaryKey"`
	Name             string    `gorm:"column:name;not null"`
	Email            string    `gorm:"column:email;not null"`
	Password         string    `gorm:"column:password;not null"`
	Photo            string    `gorm:"column:photo;not null"`
	VerificationCode string    `gorm:"column:verification_code;not null"`
	IsVerified       bool      `gorm:"column:is_verified;not null"`
	CreatedAt        time.Time `gorm:"column:created_at;not null"`
	UpdatedAt        time.Time `gorm:"column:updated_at;not null"`
}

func (impl *User) TableName() string {
	return "users"
}
