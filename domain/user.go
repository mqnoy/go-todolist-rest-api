package domain

import (
	"github.com/mqnoy/go-todolist-rest-api/dto"
	"github.com/mqnoy/go-todolist-rest-api/model"
)

type UserUseCase interface {
	RegisterUser(request dto.RegisterRequest) (*dto.User, error)
	LoginUser(payload dto.LoginRequest) (*dto.LoginResponse, error)
	GetMemberByUserId(userId string) (*model.Member, error)
	GetMemberInfo(param dto.MemberInfoParam) (*dto.User, error)
}

type UserRepository interface {
	SelectMemberByUserId(userId string) (*model.Member, error)
	InsertMember(row model.Member) (*model.Member, error)
	SelectMemberByEmail(email string) (*model.Member, error)
}
