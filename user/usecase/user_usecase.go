package usecase

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/mqnoy/go-todolist-rest-api/domain"
	"github.com/mqnoy/go-todolist-rest-api/dto"
	"github.com/mqnoy/go-todolist-rest-api/model"
	"github.com/mqnoy/go-todolist-rest-api/pkg/cerror"
	"github.com/mqnoy/go-todolist-rest-api/pkg/clogger"
	"gorm.io/gorm"
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

func (u *userUseCase) GetMemberByUserId(userId string) (*model.Member, error) {
	row, err := u.userRepo.SelectMemberByUserId(userId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, cerror.WrapError(http.StatusNotFound, fmt.Errorf("user not found"))
		}

		clogger.Logger().SetReportCaller(true)
		clogger.Logger().Errorf(err.Error())
		return nil, cerror.WrapError(http.StatusInternalServerError, fmt.Errorf("internal server error"))
	}

	return row, nil
}
