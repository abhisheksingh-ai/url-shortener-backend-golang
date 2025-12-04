package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type URL struct {
	Id          string `gorm:"column:Id;primaryKey"`
	UserId      string `gorm:"column:UserId;not null"`
	ShortCode   string `gorm:"column:ShortCode; not null"`
	OriginalUrl string `gorm:"column:OriginalUrl; not null"`
	Click       int32  `gorm:"column:Click;default:0"`

	CreatedAt time.Time `gorm:"column:CreatedAt; autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:UpdatedAt; autoUpateTime"`
}

func (URL) TableName() string {
	return "url"
}

func (u *URL) BeforeCreate(tx *gorm.DB) (err error) {
	if u.Id == "" {
		u.Id = uuid.New().String()
	}
	return
}
