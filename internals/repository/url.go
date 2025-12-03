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
	IncreaseClick(ctx context.Context, shortUrl string) error
	GetByOriginalUrl(ctx context.Context, originalUrl string, userId string) (*model.URL, error)
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
	if err := r.db.WithContext(ctx).Create(url).Error; err != nil {
		return nil, err
	}
	return url, nil
}

func (r *urlRepo) GetByShortCode(ctx context.Context, ShortUrl string) (*model.URL, error) {
	var url model.URL

	if err := r.db.WithContext(ctx).Where("ShortUrl=?", ShortUrl).First(&url).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return &url, nil
}

func (r *urlRepo) IncreaseClick(ctx context.Context, shortUrl string) error {

	var url model.URL
	if err := r.db.WithContext(ctx).First(&url, "ShortUrl=?", shortUrl).Error; err != nil {
		return err
	}

	// Increment the count by one
	if err := r.db.WithContext(ctx).Model(&url).UpdateColumn(
		"click", gorm.Expr("click + ?", 1)).Error; err != nil {
		return err
	}

	// Return the new click
	url.Click += 1

	return nil
}

func (r *urlRepo) GetByOriginalUrl(ctx context.Context, originalUrl string, userId string) (*model.URL, error) {
	var url model.URL

	if err := r.db.WithContext(ctx).Where("UserId=?", userId).Where("OriginalUrl=?", originalUrl).First(&url).Error; err != nil {
		r.logger.Error("Error: " + err.Error())
		return nil, err
	}

	return &url, nil
}
