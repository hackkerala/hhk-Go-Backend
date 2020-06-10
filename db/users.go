package db

import (
	"github.com/jinzhu/gorm"
)

//User type for anonymous user
type User struct {
	gorm.Model
	//ID    uint   `json:"id" gorm:"primary_key" `
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
}

type NewUser struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
}

// Article type
type Article struct {
	gorm.Model

	Author User   `json:"author"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}
