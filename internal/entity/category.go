package entity

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name string `gorm:"column:name;not null"`

	ProductID uint
}

func (impl *Category) TableName() string {
	return "categories"
}
