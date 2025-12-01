package repository

import (
	"context"
	"urlShortener/internals/model"
	"urlShortener/utils"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *model.User) (*model.User, error)
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
}

// This class will inherit the interface
type userRepository struct {
	db     *gorm.DB
	logger utils.Logger
}

// Constructor
func GetUserRepository(db *gorm.DB, l utils.Logger) UserRepository {
	return &userRepository{
		db:     db,
		logger: l,
	}
}

// Implementation
func (u *userRepository) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	if err := u.db.WithContext(ctx).Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// Get user by email id
func (u *userRepository) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	if err := u.db.WithContext(ctx).Where("Email=?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
