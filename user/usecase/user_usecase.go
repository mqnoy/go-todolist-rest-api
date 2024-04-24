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
	"golang.org/x/crypto/bcrypt"
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
func (u *userUseCase) RegisterUser(request dto.RegisterRequest) (*dto.User, error) {

	// Generate hashed password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		clogger.Logger().SetReportCaller(true)
		clogger.Logger().Errorf(err.Error())
		return nil, cerror.WrapError(http.StatusInternalServerError, fmt.Errorf("internal server error"))
	}

	row := model.Member{
		User: model.User{
			FullName: request.FullName,
			Email:    request.Email,
			Password: string(hashedPassword),
		},
	}

	member, err := u.userRepo.InsertMember(row)
	if err != nil {
		clogger.Logger().SetReportCaller(true)
		clogger.Logger().Errorf(err.Error())
		return nil, cerror.WrapError(http.StatusInternalServerError, fmt.Errorf("internal server error"))
	}

	return u.ComposeUser(member), nil
}

func (u *userUseCase) ComposeUser(m *model.Member) *dto.User {
	return &dto.User{
		ID:       m.User.ID,
		FullName: m.User.FullName,
		Email:    m.User.Email,
		Member: dto.Member{
			ID: m.ID,
		},
		Timestamp: dto.ComposeTimestamp(m.User.TimestampColumn),
	}
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
