package dto

import "context"

const AuthorizedUserCtxKey = "authorizedUser"

type AuthorizedUser struct {
	UserID     string
	MemberID   string
	Privileges []string
}

type MemberInfoParam struct {
	SubjectId string
}

func GetAuthorizedUser(ctx context.Context) AuthorizedUser {
	au := ctx.Value(AuthorizedUserCtxKey).(AuthorizedUser)
	return AuthorizedUser{
		UserID:   au.UserID,
		MemberID: au.MemberID,
	}
}
