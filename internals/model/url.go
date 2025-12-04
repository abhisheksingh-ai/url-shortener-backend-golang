package model

import (
	"time"
)

type URL struct {
	UserId      string    `gorm:"column:UserId"`
	ShortCode   string    `gorm:"column:ShortCode;primaryKey"`
	OriginalUrl string    `gorm:"column:OriginalUrl"`
	CreatedAt   time.Time `gorm:"column:CreatedAt"`
	Click       int32     `gorm:"column:Click"`
}

func (URL) TableName() string {
	return "url"
}
