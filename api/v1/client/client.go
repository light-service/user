package client

import (
	"context"
	"time"
)

type UserInfoRegisterView struct {
	Username  string
	Password  string
	FirstName string
	LastName  string
	Phone     string
	Email     string
}

type UserInfoModifyView struct {
	FirstName string
	LastName  string
	Phone     string
	Email     string
}

type UserInfo struct {
	ID        int
	Username  string
	FirstName string
	LastName  string
	Phone     string
	Email     string
	LastLogin time.Time
}

type UserServiceClient interface {
	Register(ctx context.Context, info *UserInfoRegisterView) (userID int, err error)
	Login(ctx context.Context, username string, password string, ip string, expireAt *time.Time) (token string, actualExpireAt time.Time, err error)
	Authenticate(ctx context.Context, token string) (userID int, err error)

	GetUserInfo(ctx context.Context, userID int) (info *UserInfo, err error)
	ModifyUserInfo(ctx context.Context, info *UserInfoModifyView, userID int) error
	ModifyPassword(ctx context.Context, newPassword string, oldPassword string, userID int) error
}
