package service

import (
	"context"
	"urlShortener/internals/dto"
	"urlShortener/internals/model"
	"urlShortener/internals/repository"
	"urlShortener/utils"
)

// interface
type UrlService interface {
	CreateNewShortUrl(ctx context.Context, urlDto *dto.UrlDto) (*dto.UrlResponseDto, error)
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

	//1. dto to model
	url := &model.URL{
		OriginalUrl: urlDto.OriginalUrl,
	}

	//2. Generate 6 digit Unique char
	var shortCode string
	for {
		shortCode = utils.GenerateShortCode(6)
		shortUrl := "http//localhost:8080/" + shortCode
		exists, _ := s.repo.GetByShortCode(ctx, shortUrl)

		if exists == nil {
			break
		}
	}

	url.ShortUrl = "http//localhost:8080/" + shortCode

	//3. Save to db via repository
	result, err := s.repo.CreateNewShortUrl(ctx, url)
	if err != nil {
		s.logger.Error("Error in creating new url: " + err.Error())
		return nil, err
	}

	// 4. model   ----> response dto
	response := &dto.UrlResponseDto{
		ShortUrl: result.ShortUrl,
		Message:  "ShortUrl Created Successfully",
		Data:     result,
	}

	return response, nil
}
