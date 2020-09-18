package rpc

import (
	"github.com/light-service/user/api/v1"
	"github.com/light-service/user/model"
)

func adapterProtoUserInfo(info *model.User) *api.UserInfo {
	if info == nil {
		return nil
	}

	return &api.UserInfo{
		Id:        int64(info.ID),
		Username:  info.Username,
		FirstName: info.FirstName,
		LastName:  info.LastName,
		Phone:     info.Phone,
		Email:     info.Email,
		LastLogin: api.AdapterProtoTime(&info.LastLoginAt),
	}
}
