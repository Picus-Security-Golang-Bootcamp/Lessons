package models

import (
	"gorm.io/gorm"
)

type Author struct {
	gorm.Model
	Name  *string `gorm:"unique"`
	Books []Book
}

func (Author) TableName() string {
	//default table name
	return "authors"
}
