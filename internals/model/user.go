package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	UserID    string `gorm:"column:UserID"`
	FirstName string `gorm:"column:FirstName"`
	LastName  string `gorm:"column:LastName"`
	Email     string `gorm:"column:Email"`
	Password  string `gorm:"column:Password"`

	CreatedAt time.Time `gorm:"column:CreatedAt"`
	UpdatedAt time.Time `gorm:"column:UpdatedAt"`
}

// Table Name
func (User) TableName() string {
	return "User"
}

// this will create the UserID
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.UserID == "" {
		u.UserID = uuid.New().String()
	}
	return
}
