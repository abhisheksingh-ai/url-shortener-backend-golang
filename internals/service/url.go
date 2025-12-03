package service

import (
	"context"
	"errors"
	"urlShortener/internals/dto"
	"urlShortener/internals/model"
	"urlShortener/internals/repository"
	"urlShortener/utils"

	"gorm.io/gorm"
)

// interface
type UrlService interface {
	CreateNewShortUrl(ctx context.Context, urlDto *dto.UrlDto) (*dto.UrlResponseDto, error)
	RedirectUrl(ctx context.Context, urlDto *dto.UrlDto) (*dto.UrlResponseDto, error)
}

// Implementation
type urlService struct {
	logger utils.Logger
	repo   repository.UrlRepo
}

// constructor
func GetUrlService(l utils.Logger, r repository.UrlRepo) UrlService {
	return &urlService{
		logger: l,
		repo:   r,
	}
}

func (s *urlService) CreateNewShortUrl(ctx context.Context, urlDto *dto.UrlDto) (*dto.UrlResponseDto, error) {

	//1. dto → model
	url := &model.URL{
		OriginalUrl: urlDto.OriginalUrl,
		UserId:      urlDto.UserId,
	}

	// 2. Check if short URL already exists for this user + originalUrl
	existingUrl, err := s.repo.GetByOriginalUrl(ctx, url.OriginalUrl, url.UserId)
	if err == nil {
		// means it exists
		return &dto.UrlResponseDto{
			ShortUrl: existingUrl.ShortUrl,
			Message:  "ShortUrl Already Exists",
			Data:     existingUrl,
		}, nil
	}
	// if err != nil and it's NOT record not found → it's a real DB error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	// 3. Generate a new unique short code
	var shortCode string
	for {
		shortCode = utils.GenerateShortCode(6)
		shortUrl := "http://localhost:8080/" + shortCode

		// check if this short code already exists
		u, _ := s.repo.GetByShortCode(ctx, shortUrl)
		if u == nil {
			break
		}
	}

	url.ShortUrl = "http://localhost:8080/" + shortCode

	//4. Save
	result, err := s.repo.CreateNewShortUrl(ctx, url)
	if err != nil {
		s.logger.Error("Error in creating new url: " + err.Error())
		return nil, err
	}

	//5. Prepare response
	response := &dto.UrlResponseDto{
		ShortUrl: result.ShortUrl,
		Message:  "ShortUrl Created Successfully",
		Data:     result,
	}

	return response, nil
}

func (s *urlService) RedirectUrl(ctx context.Context, urlDto *dto.UrlDto) (*dto.UrlResponseDto, error) {
	//1. dto ---> getting data
	shortCode := urlDto.ShortCode

	// Making the short url
	shortUrl := "http://localhost:8080/" + shortCode

	//Checking this in database
	exists, _ := s.repo.GetByShortCode(ctx, shortUrl)

	if exists == nil {
		return &dto.UrlResponseDto{
			Message: "URL Not Found",
			Data:    nil,
		}, nil
	}

	// Increase the count by one
	err := s.repo.IncreaseClick(ctx, exists.ShortUrl)
	if err != nil {
		return &dto.UrlResponseDto{
			Message: "Issue in increment the click column",
		}, err
	}

	return &dto.UrlResponseDto{
		OriginalUrl: exists.OriginalUrl,
	}, nil
}
