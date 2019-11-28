package db

import "github.com/jinzhu/gorm"

type Account struct {
	gorm.Model
	Account       string // 10 bytes
	Password      string // 10 bytes
	Name          string // 20 bytes
	Idcard        string // 14 bytes
	Phone         string // 14 bytes
	Question1     string // 20 bytes
	Answer1       string // 12 bytes
	Email         string // 51 bytes
	Question2     string // 20 bytes
	Answer2       string // 12 bytes
	Birthday      string // 10 bytes
	Mobilephone   string // 55 bytes
	Locked        bool
	Wrongpwdtimes int
	Online        bool
	Session       string
}
