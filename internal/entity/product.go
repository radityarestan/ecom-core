package entity

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Stock        uint    `gorm:"column:stock;not null"`
	Sold         uint    `gorm:"column:sold;not null"`
	AmountRating uint    `gorm:"column:amount_rating;not null"`
	Name         string  `gorm:"column:name;not null"`
	Detail       string  `gorm:"column:detail;not null"`
	Photo        string  `gorm:"column:photo;not null"`
	Rating       float32 `gorm:"column:rating;not null"`
	Price        float64 `gorm:"column:price;not null"`

	UserID   uint
	Category Category
}

func (impl *Product) TableName() string {
	return "products"
}
