package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	UserID    string `gorm:"column:userid"`
	FirstName string `gorm:"column:firstname"`
	LastName  string `gorm:"column:lastname"`
	Email     string `gorm:"column:email"`
	Password  string `gorm:"column:password"`

	CreatedAt time.Time `gorm:"column:createdat"`
	UpdatedAt time.Time `gorm:"column:updatedat"`
}

// Table Name
func (User) TableName() string {
	return "user"
}

// this will create the UserID
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.UserID == "" {
		u.UserID = uuid.New().String()
	}
	return
}
