package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserId   string
	Username string
	Password string
	Email    string `gorm:"uniqueIndex"`
	GithubId string
	Sessions []Session
}
