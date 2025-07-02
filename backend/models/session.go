package models

import "gorm.io/gorm"

type Session struct {
	gorm.Model
	RepoURL string
	Path    string
	UserID  string
}
