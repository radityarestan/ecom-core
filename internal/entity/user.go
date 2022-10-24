package entity

import "time"

type User struct {
	ID        int64     `gorm:"primaryKey"`
	Name      string    `gorm:"column:name"`
	CreatedAt time.Time `gorm:"column:created_at"`
}

func (impl *User) TableName() string {
	return "users"
}
