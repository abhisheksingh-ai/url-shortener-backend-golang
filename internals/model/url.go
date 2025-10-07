package model

import (
	"time"
)

type URL struct {
	ShortUrl    string    `gorm:"column:shorturl;primaryKey"`
	OriginalUrl string    `gorm:"column:originalurl"`
	CreatedAt   time.Time `gorm:"column:createdat"`
}

func (URL) TableName() string {
	return "url"
}
