package client

import (
	"context"
	"time"

	"github.com/light-service/user/api/v1"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

type userServiceClientImpl struct {
	client api.UserServiceClient
}

func NewUserServiceClient(ctx context.Context, serviceAddr string, opts ...grpc.DialOption) UserServiceClient {
	if opts == nil {
		opts = []grpc.DialOption{
			grpc.WithInsecure(),
			grpc.WithKeepaliveParams(keepalive.ClientParameters{
				Time:                10 * time.Second,
				Timeout:             10 * time.Second,
				PermitWithoutStream: true,
			}),
		}
	}

	conn, err := grpc.DialContext(ctx, serviceAddr, opts...)
	if err != nil {
		log.WithError(err).Fatalf("dial user service error: addr=%s", serviceAddr)
	}

	client := api.NewUserServiceClient(conn)
	return &userServiceClientImpl{client}
}

func (u *userServiceClientImpl) Register(ctx context.Context, info *UserInfoRegisterView) (userID int, err error) {
	req := &api.RegisterRequest{
		Username:  info.Username,
		Password:  info.Password,
		FirstName: info.FirstName,
		LastName:  info.LastName,
		Phone:     info.Phone,
		Email:     info.Email,
	}
	resp, err := u.client.Register(ctx, req)
	if err != nil {
		return 0, err
	}
	return int(resp.GetUserId()), nil
}

func (u *userServiceClientImpl) Login(ctx context.Context, username string, password string, ip string, expireAt *time.Time) (token string, actualExpireAt time.Time, err error) {
	req := &api.LoginRequest{
		Username: username,
		Password: password,
		Ip:       ip,
		ExpireAt: api.AdapterProtoTime(expireAt),
	}
	resp, err := u.client.Login(ctx, req)
	if err != nil {
		return "", time.Time{}, err
	}

	return resp.GetToken(), api.Time(api.AdapterTime(resp.GetExpireAt())), nil
}

func (u *userServiceClientImpl) Authenticate(ctx context.Context, token string) (userID int, err error) {
	req := &api.AuthenticateRequest{
		Token: token,
	}
	resp, err := u.client.Authenticate(ctx, req)
	if err != nil {
		return 0, err
	}
	return int(resp.GetUserId()), nil
}

func (u *userServiceClientImpl) GetUserInfo(ctx context.Context, userID int) (info *UserInfo, err error) {
	req := &api.GetUserInfoRequest{
		UserId: int64(userID),
	}
	userInfo, err := u.client.GetUserInfo(ctx, req)
	if err != nil {
		return
	}
	return adapterUserInfo(userInfo), nil
}

func (u *userServiceClientImpl) ModifyUserInfo(ctx context.Context, info *UserInfoModifyView, userID int) error {
	req := &api.ModifyUserInfoRequest{
		FirstName: info.FirstName,
		LastName:  info.LastName,
		Phone:     info.Phone,
		Email:     info.Email,
		UserId:    int64(userID),
	}
	_, err := u.client.ModifyUserInfo(ctx, req)
	return err
}

func (u *userServiceClientImpl) ModifyPassword(ctx context.Context, newPassword string, oldPassword string, userID int) error {
	req := &api.ModifyPasswordRequest{
		NewPassword: newPassword,
		OldPassword: oldPassword,
		UserId:      int64(userID),
	}
	_, err := u.client.ModifyPassword(ctx, req)
	return err
}
