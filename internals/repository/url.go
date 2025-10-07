package repository

import (
	"context"
	"urlShortener/internals/model"
	"urlShortener/utils"

	"gorm.io/gorm"
)

// Interface
type UrlRepo interface {
	CreateNewShortUrl(ctx context.Context, url *model.URL) (*model.URL, error)
	GetByShortCode(ctx context.Context, ShortUrl string) (*model.URL, error)
}

// Implementing UrlRepo
type urlRepo struct {
	logger utils.Logger
	db     *gorm.DB
}

// constructor that will return UrlRepo
func GetUrlRepo(l utils.Logger, db *gorm.DB) UrlRepo {
	return &urlRepo{
		logger: l,
		db:     db,
	}
}

func (r *urlRepo) CreateNewShortUrl(ctx context.Context, url *model.URL) (*model.URL, error) {
	if err := r.db.WithContext(ctx).Create(&url).Error; err != nil {
		return nil, err
	}
	return url, nil
}

func (r *urlRepo) GetByShortCode(ctx context.Context, ShortUrl string) (*model.URL, error) {
	var url model.URL

	if err := r.db.WithContext(ctx).Where(`"shorturl" = ?`, ShortUrl).First(&url).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return &url, nil
}
