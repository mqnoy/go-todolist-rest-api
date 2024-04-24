package usecase

import (
	"github.com/mqnoy/go-todolist-rest-api/domain"
	"github.com/mqnoy/go-todolist-rest-api/dto"
)

type userUseCase struct {
	userRepo domain.UserRepository
}

func New(userRepo domain.UserRepository) domain.UserUseCase {
	return &userUseCase{
		userRepo: userRepo,
	}
}

// LoginUser implements domain.UserUseCase.
func (u *userUseCase) LoginUser(payload *dto.LoginRequest) (*dto.LoginResponse, error) {
	return &dto.LoginResponse{
		AccessToken:  "",
		RefreshToken: "",
		User:         dto.User{},
	}, nil
}

// RegisterUser implements domain.UserUseCase.
func (u *userUseCase) RegisterUser(payload *dto.RegisterRequest) (*dto.RegisterResponse, error) {
	return &dto.RegisterResponse{}, nil
}
