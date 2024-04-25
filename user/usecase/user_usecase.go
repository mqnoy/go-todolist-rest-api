package usecase

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/mqnoy/go-todolist-rest-api/config"
	"github.com/mqnoy/go-todolist-rest-api/domain"
	"github.com/mqnoy/go-todolist-rest-api/dto"
	"github.com/mqnoy/go-todolist-rest-api/model"
	"github.com/mqnoy/go-todolist-rest-api/pkg/cerror"
	"github.com/mqnoy/go-todolist-rest-api/pkg/clogger"
	"github.com/mqnoy/go-todolist-rest-api/pkg/token"
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
func (u *userUseCase) LoginUser(payload dto.LoginRequest) (*dto.LoginResponse, error) {
	memberRow, err := u.GetUserByEmail(payload.Email)
	if err != nil {
		return nil, err
	}

	userRow := memberRow.User
	if userRow.Email == "" {
		return nil, cerror.WrapError(http.StatusBadRequest, fmt.Errorf("email not found"))
	}

	// Compare password
	if err := u.ComparePassword(userRow.Password, payload.Password); err != nil {
		return nil, err
	}

	// Generate accessToken
	accessTknExpiry := jwt.NewNumericDate(time.Now().Add(time.Duration(config.AppConfig.JWT.AccessTokenExpiry) * time.Second))
	accessTkn, err := u.GenerateToken(accessTknExpiry, userRow.ID)
	if err != nil {
		return nil, err
	}

	// Generate refreshToken
	refreshTknExpiry := jwt.NewNumericDate(time.Now().Add(time.Duration(config.AppConfig.JWT.RefreshTokenExpiry) * time.Second))
	refreshTkn, err := u.GenerateToken(refreshTknExpiry, userRow.ID)
	if err != nil {
		return nil, err
	}

	return &dto.LoginResponse{
		AccessToken:  accessTkn,
		RefreshToken: refreshTkn,
		User:         u.ComposeUserMember(memberRow),
	}, nil
}

func (u *userUseCase) GenerateToken(expiredIn *jwt.NumericDate, subjectId string) (string, error) {
	key := []byte(config.AppConfig.JWT.Key)
	mapClaims := token.GenerateMapClaims(token.CustomClaimOptions{
		ExpiredTime: expiredIn,
		SubjectId:   subjectId,
	})

	token, err := token.Generate(mapClaims, key)
	if err != nil {
		return "", cerror.WrapError(http.StatusInternalServerError, err)
	}

	return token, nil
}

func (u *userUseCase) ComparePassword(password string, inputPassword string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(password), []byte(inputPassword)); err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return cerror.WrapError(http.StatusInternalServerError, fmt.Errorf("password doesn't match"))
		}

		clogger.Logger().SetReportCaller(true)
		clogger.Logger().Errorf(err.Error())
		return cerror.WrapError(http.StatusInternalServerError, fmt.Errorf("internal server error"))
	}

	return nil
}

// RegisterUser implements domain.UserUseCase.
func (u *userUseCase) RegisterUser(request dto.RegisterRequest) (*dto.User, error) {

	// Validate email is exist
	existEmail, err := u.GetUserByEmail(request.Email)
	if existEmail != nil && err == nil {
		return nil, cerror.WrapError(http.StatusBadRequest, fmt.Errorf("email already exist"))
	}

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

	return u.ComposeUserMember(member), nil
}

func (u *userUseCase) ComposeUserMember(m *model.Member) *dto.User {
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

func (u *userUseCase) ComposeUser(m *model.User) *dto.User {
	return &dto.User{
		ID:       m.ID,
		FullName: m.FullName,
		Email:    m.Email,
		Member: dto.Member{
			ID: m.ID,
		},
		Timestamp: dto.ComposeTimestamp(m.TimestampColumn),
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

func (u *userUseCase) GetUserByEmail(email string) (*model.Member, error) {
	row, err := u.userRepo.SelectMemberByEmail(email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("member not found")
		}

		clogger.Logger().SetReportCaller(true)
		clogger.Logger().Errorf(err.Error())
		return nil, cerror.WrapError(http.StatusInternalServerError, fmt.Errorf("internal server error"))
	}

	return row, nil
}

func (u *userUseCase) GetMemberInfo(param dto.MemberInfoParam) (*dto.User, error) {
	row, err := u.userRepo.SelectMemberByUserId(param.SubjectId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("member not found")
		}

		clogger.Logger().SetReportCaller(true)
		clogger.Logger().Errorf(err.Error())
		return nil, cerror.WrapError(http.StatusInternalServerError, fmt.Errorf("internal server error"))
	}

	return u.ComposeUserMember(row), nil
}
