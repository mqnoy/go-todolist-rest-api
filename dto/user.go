package dto

import "net/http"

type User struct{}

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

type RegisterRequest struct{}

type RegisterResponse struct{}
