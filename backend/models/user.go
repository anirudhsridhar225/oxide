package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID            string `gorm:"type:uuid;primaryKey"`
	Username      string
	Password      string
	Email         string `gorm:"uniqueIndex"`
	GithubId      string
	Sessions      []Session
	EmailVerified bool
	CreatedAt     time.Time
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.NewString()
	u.EmailVerified = false
	u.CreatedAt = time.Now()
	return
}
