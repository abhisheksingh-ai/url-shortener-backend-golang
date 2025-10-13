package service

import (
	"context"
	"urlShortener/internals/dto"
	"urlShortener/internals/model"
	"urlShortener/internals/repository"
	"urlShortener/utils"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	CreateUser(ctx context.Context, userDto *dto.UserDto) (*dto.UserResponseDto, error)
}

// inheritence
type userService struct {
	userRepo repository.UserRepository
	logger   utils.Logger
}

// constructor
func GetNewService(r repository.UserRepository, l utils.Logger) UserService {
	return &userService{
		userRepo: r,
		logger:   l,
	}
}

// Create new user service
func (s *userService) CreateUser(ctx context.Context, userDto *dto.UserDto) (*dto.UserResponseDto, error) {
	//1. userDto to model
	user := &model.User{
		FirstName: userDto.FirstName,
		LastName:  userDto.LastName,
		Email:     userDto.Email,
		Password:  userDto.Password,
	}

	//2. hash this password to store it
	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		s.logger.Error("Error in password hashing" + err.Error())
		return &dto.UserResponseDto{
			Message: "Error in password hashing",
		}, err
	}

	user.Password = string(hashed)

	//3. Call the service
	data, err := s.userRepo.CreateUser(ctx, user)
	if err != nil {
		s.logger.Error("Error in user creation" + err.Error())
	}

	//4. Generate token after user signup
	token, err := utils.GenerateToken(data.UserID, data.Email)
	if err != nil {
		s.logger.Error("Error in generating token: " + err.Error())
		return nil, err
	}

	return &dto.UserResponseDto{
		Message: "User created succesfully",
		UserID:  data.UserID,
		Token:   token,
	}, nil
}
