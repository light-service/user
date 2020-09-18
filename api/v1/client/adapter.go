package client

import (
	"github.com/light-service/user/api/v1"
)

func adapterUserInfo(info *api.UserInfo) *UserInfo {
	if info == nil {
		return nil
	}

	return &UserInfo{
		ID:        int(info.Id),
		Username:  info.Username,
		FirstName: info.FirstName,
		LastName:  info.LastName,
		Phone:     info.Phone,
		Email:     info.Email,
		LastLogin: api.Time(api.AdapterTime(info.LastLogin)),
	}
}

