package dto

import "urlShortener/internals/model"

type UrlDto struct {
	OriginalUrl string `json:"originalUrl"`
	ShortUrl    string `json:"shortUrl"`
}

type UrlResponseDto struct {
	ShortUrl    string     `json:"shortUrl,omitempty"`
	Status      string     `json:"status,omitempty"`
	Message     string     `json:"message,omitempty"`
	Data        *model.URL `json:"data,omitempty"`
	OriginalUrl string     `json:"originalUrl,omitempty"`
}
