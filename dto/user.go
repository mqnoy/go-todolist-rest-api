package dto

import "net/http"

type User struct {
	ID       string `json:"id"`
	FullName string `json:"fullName"`
	Email    string `json:"email"`
	Member   Member `json:"member"`
	Timestamp
}

type Member struct {
	ID string `json:"id"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (l *LoginRequest) Bind(r *http.Request) error {
	return nil
}

type LoginResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	User         User   `json:"user"`
}

type RegisterRequest struct {
	FullName string `json:"fullName" validate:"required"`
	Email    string `json:"email" validate:"email,normalize"`
	Password string `json:"password" validate:"required"`
}

func (rr *RegisterRequest) Bind(r *http.Request) error {
	return nil
}

type RegisterResponse struct{}
