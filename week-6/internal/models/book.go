package models

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Name        *string `gorm:"unique"`
	Genre       string
	Isbn        string
	PageNumber  int64
	ReleaseDate string
	AuthorID    uint
	Author      Author
}

func (Book) TableName() string {
	//default table name
	return "books"
}
