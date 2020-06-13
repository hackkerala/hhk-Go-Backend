package db

import (
	"regexp"

	"github.com/jinzhu/gorm"
)

var rxEmail = regexp.MustCompile(".+@.+\\..+")

//User type for anonymous user
type User struct {
	gorm.Model
	//ID    uint   `json:"id" gorm:"primary_key" `
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
}

func (u *User) Emailval() bool {

	match := rxEmail.Match([]byte(u.Email))

	if match == false {
		return false
	}
	return true
}
