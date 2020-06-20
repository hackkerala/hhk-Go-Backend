package db

import (
	"regexp"

	"github.com/jinzhu/gorm"
)

type Hackathon struct {
	gorm.Model

	Name       string `json:"name"`
	Organizers []User `json:"organizers"`
	Place      string `json:"place"`
	Date       string `json:"date"`
	URL        string `json:"url"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
}

var phoneReg = regexp.MustCompile("^((\\+91|91|0)[\\- ]{0,1})?[456789]\\d{9}$")
var urlreg = regexp.MustCompile(`[(http(s)?):\/\/(www\.)?a-zA-Z0-9@:%._\+~#=]{2,256}\.[a-z]{2,6}\b([-a-zA-Z0-9@:%_\+.~#?&//=]*)`)

func (h *Hackathon) Emaileval() bool {

	em := rxEmail.Match([]byte(h.Email))

	if em == false {
		return false
	}
	return true
}
func (h *Hackathon) Phoneeval() bool {
	ph := phoneReg.Match([]byte(h.Phone))

	if ph == false {
		return false
	}
	return true
}
func (h *Hackathon) Urleval() bool {
	url := urlreg.Match([]byte(h.URL))
	if url == false {
		return false
	}
	return true
}
