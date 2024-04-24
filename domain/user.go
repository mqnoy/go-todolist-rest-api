package domain

import "github.com/mqnoy/go-todolist-rest-api/dto"

type UserUseCase interface {
	RegisterUser(payload *dto.RegisterRequest) (*dto.RegisterResponse, error)
	LoginUser(payload *dto.LoginRequest) (*dto.LoginResponse, error)
}

type UserRepository interface{}
