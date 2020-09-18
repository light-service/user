package usecase

import (
	"github.com/light-service/user/model"
	"time"
)

type UseCase interface {
	Register(username string, password string, firstName string, lastName string, phone string, email string) (userID int, err error)
	Login(username string, password string, ip string, expireAt *time.Time) (token string, actualExpireAt time.Time, err error)
	Authenticate(token string) (userID int, err error)

	GetUserInfo(userID int) (info *model.User, err error)
	ModifyUserInfo(firstName string, lastName string, phone string, email string, userID int) error
	ModifyPassword(newPassword string, oldPassword string, userID int) error
}
