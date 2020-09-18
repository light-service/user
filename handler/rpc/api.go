package rpc

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/light-service/user/api/v1"
	"github.com/light-service/user/model/usecase"
	"google.golang.org/grpc"
)

func Serve(grpcServer *grpc.Server, ucase usecase.UseCase) {
	h := &handler{ucase}
	api.RegisterUserServiceService(grpcServer, &api.UserServiceService{
		Register:       h.Register,
		Login:          h.Login,
		Authenticate:   h.Authenticate,
		GetUserInfo:    h.GetUserInfo,
		ModifyUserInfo: h.ModifyUserInfo,
		ModifyPassword: h.ModifyPassword,
	})
}

type handler struct {
	ucase usecase.UseCase
}

func (h *handler) Register(context context.Context, request *api.RegisterRequest) (*api.RegisterResponse, error) {
	userId, err := h.ucase.Register(request.GetUsername(), request.GetPassword(), request.GetFirstName(), request.GetLastName(), request.GetPhone(), request.GetEmail())
	if err != nil {
		return nil, err
	}

	return &api.RegisterResponse{
		UserId: int64(userId),
	}, nil
}

func (h *handler) Login(context context.Context, request *api.LoginRequest) (*api.LoginResponse, error) {
	token, expireAt, err := h.ucase.Login(request.GetUsername(), request.GetPassword(), request.GetIp(), api.AdapterTime(request.GetExpireAt()))
	if err != nil {
		return nil, err
	}

	return &api.LoginResponse{
		Token:    token,
		ExpireAt: api.AdapterProtoTime(&expireAt),
	}, nil
}

func (h *handler) Authenticate(context context.Context, request *api.AuthenticateRequest) (*api.AuthenticateResponse, error) {
	userID, err := h.ucase.Authenticate(request.GetToken())
	if err != nil {
		return nil, err
	}

	return &api.AuthenticateResponse{
		UserId: int64(userID),
	}, nil
}

func (h *handler) GetUserInfo(context context.Context, request *api.GetUserInfoRequest) (*api.UserInfo, error) {
	userInfo, err := h.ucase.GetUserInfo(int(request.GetUserId()))
	if err != nil {
		return nil, err
	}

	return adapterProtoUserInfo(userInfo), nil
}

func (h *handler) ModifyUserInfo(context context.Context, request *api.ModifyUserInfoRequest) (*empty.Empty, error) {
	err := h.ucase.ModifyUserInfo(request.GetFirstName(), request.GetLastName(), request.GetPhone(), request.GetEmail(), int(request.GetUserId()))
	if err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}

func (h *handler) ModifyPassword(context context.Context, request *api.ModifyPasswordRequest) (*empty.Empty, error) {
	err := h.ucase.ModifyPassword(request.GetNewPassword(), request.GetOldPassword(), int(request.GetUserId()))
	if err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}
