package model

import (
	"time"
)

type URL struct {
	UserId      string    `gorm:"column:UserId"`
	ShortUrl    string    `gorm:"column:ShortUrl;primaryKey"`
	OriginalUrl string    `gorm:"column:OriginalUrl"`
	CreatedAt   time.Time `gorm:"column:CreatedAt"`
	Click       int32     `gorm:"column:Click"`
}

func (URL) TableName() string {
	return "url"
}
