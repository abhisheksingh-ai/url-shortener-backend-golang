package service

import (
	"context"
	"log/slog"
	"urlShortener/internals/repository"
	"urlShortener/utils"

	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Login(ctx context.Context, email, password string) (string, error)
}

type authService struct {
	userRepo repository.UserRepository
	logger   *slog.Logger
}

// Authservice object creation
func GetAuthService(r repository.UserRepository, l *slog.Logger) AuthService {
	return &authService{
		userRepo: r,
		logger:   l,
	}
}

func (s *authService) Login(ctx context.Context, email, password string) (string, error) {
	//1. fetch user by email
	user, err := s.userRepo.GetUserByEmail(ctx, email)
	if err != nil {
		s.logger.Error("Error in finding the user by email: " + err.Error())
		return "", err
	}

	// 2. compare password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", err
	}

	// generate jwt
	token, err := utils.GenerateToken(user.UserID, user.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}
