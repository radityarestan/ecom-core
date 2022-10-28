package entity

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name string `gorm:"column:name;not null"`

	Products []Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (impl *Category) TableName() string {
	return "categories"
}
