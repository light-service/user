// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

/*
Package api is a generated protocol buffer package.

It is generated from these files:
	user.proto

It has these top-level messages:
	UserInfo
	RegisterRequest
	RegisterResponse
	LoginRequest
	LoginResponse
	AuthenticateRequest
	AuthenticateResponse
	GetUserInfoRequest
	ModifyUserInfoRequest
	ModifyPasswordRequest
*/
package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"
import google_protobuf1 "github.com/golang/protobuf/ptypes/empty"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type UserInfo struct {
	Id        int64                      `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Username  string                     `protobuf:"bytes,2,opt,name=username" json:"username,omitempty"`
	FirstName string                     `protobuf:"bytes,3,opt,name=first_name,json=firstName" json:"first_name,omitempty"`
	LastName  string                     `protobuf:"bytes,4,opt,name=last_name,json=lastName" json:"last_name,omitempty"`
	Phone     string                     `protobuf:"bytes,5,opt,name=phone" json:"phone,omitempty"`
	Email     string                     `protobuf:"bytes,6,opt,name=email" json:"email,omitempty"`
	LastLogin *google_protobuf.Timestamp `protobuf:"bytes,7,opt,name=last_login,json=lastLogin" json:"last_login,omitempty"`
}

func (m *UserInfo) Reset()                    { *m = UserInfo{} }
func (m *UserInfo) String() string            { return proto.CompactTextString(m) }
func (*UserInfo) ProtoMessage()               {}
func (*UserInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *UserInfo) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UserInfo) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *UserInfo) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *UserInfo) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *UserInfo) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *UserInfo) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UserInfo) GetLastLogin() *google_protobuf.Timestamp {
	if m != nil {
		return m.LastLogin
	}
	return nil
}

type RegisterRequest struct {
	Username  string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	Password  string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
	FirstName string `protobuf:"bytes,3,opt,name=first_name,json=firstName" json:"first_name,omitempty"`
	LastName  string `protobuf:"bytes,4,opt,name=last_name,json=lastName" json:"last_name,omitempty"`
	Phone     string `protobuf:"bytes,5,opt,name=phone" json:"phone,omitempty"`
	Email     string `protobuf:"bytes,6,opt,name=email" json:"email,omitempty"`
}

func (m *RegisterRequest) Reset()                    { *m = RegisterRequest{} }
func (m *RegisterRequest) String() string            { return proto.CompactTextString(m) }
func (*RegisterRequest) ProtoMessage()               {}
func (*RegisterRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *RegisterRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *RegisterRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *RegisterRequest) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *RegisterRequest) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *RegisterRequest) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *RegisterRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type RegisterResponse struct {
	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
}

func (m *RegisterResponse) Reset()                    { *m = RegisterResponse{} }
func (m *RegisterResponse) String() string            { return proto.CompactTextString(m) }
func (*RegisterResponse) ProtoMessage()               {}
func (*RegisterResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *RegisterResponse) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type LoginRequest struct {
	Username string                     `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	Password string                     `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
	Ip       string                     `protobuf:"bytes,3,opt,name=ip" json:"ip,omitempty"`
	ExpireAt *google_protobuf.Timestamp `protobuf:"bytes,4,opt,name=expire_at,json=expireAt" json:"expire_at,omitempty"`
}

func (m *LoginRequest) Reset()                    { *m = LoginRequest{} }
func (m *LoginRequest) String() string            { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()               {}
func (*LoginRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *LoginRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *LoginRequest) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *LoginRequest) GetExpireAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.ExpireAt
	}
	return nil
}

type LoginResponse struct {
	Token    string                     `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
	ExpireAt *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=expire_at,json=expireAt" json:"expire_at,omitempty"`
}

func (m *LoginResponse) Reset()                    { *m = LoginResponse{} }
func (m *LoginResponse) String() string            { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()               {}
func (*LoginResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *LoginResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *LoginResponse) GetExpireAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.ExpireAt
	}
	return nil
}

type AuthenticateRequest struct {
	Token string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
}

func (m *AuthenticateRequest) Reset()                    { *m = AuthenticateRequest{} }
func (m *AuthenticateRequest) String() string            { return proto.CompactTextString(m) }
func (*AuthenticateRequest) ProtoMessage()               {}
func (*AuthenticateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *AuthenticateRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type AuthenticateResponse struct {
	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
}

func (m *AuthenticateResponse) Reset()                    { *m = AuthenticateResponse{} }
func (m *AuthenticateResponse) String() string            { return proto.CompactTextString(m) }
func (*AuthenticateResponse) ProtoMessage()               {}
func (*AuthenticateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *AuthenticateResponse) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type GetUserInfoRequest struct {
	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
}

func (m *GetUserInfoRequest) Reset()                    { *m = GetUserInfoRequest{} }
func (m *GetUserInfoRequest) String() string            { return proto.CompactTextString(m) }
func (*GetUserInfoRequest) ProtoMessage()               {}
func (*GetUserInfoRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *GetUserInfoRequest) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type ModifyUserInfoRequest struct {
	FirstName string `protobuf:"bytes,1,opt,name=first_name,json=firstName" json:"first_name,omitempty"`
	LastName  string `protobuf:"bytes,2,opt,name=last_name,json=lastName" json:"last_name,omitempty"`
	Phone     string `protobuf:"bytes,3,opt,name=phone" json:"phone,omitempty"`
	Email     string `protobuf:"bytes,4,opt,name=email" json:"email,omitempty"`
	UserId    int64  `protobuf:"varint,5,opt,name=user_id,json=userId" json:"user_id,omitempty"`
}

func (m *ModifyUserInfoRequest) Reset()                    { *m = ModifyUserInfoRequest{} }
func (m *ModifyUserInfoRequest) String() string            { return proto.CompactTextString(m) }
func (*ModifyUserInfoRequest) ProtoMessage()               {}
func (*ModifyUserInfoRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *ModifyUserInfoRequest) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *ModifyUserInfoRequest) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *ModifyUserInfoRequest) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *ModifyUserInfoRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *ModifyUserInfoRequest) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type ModifyPasswordRequest struct {
	NewPassword string `protobuf:"bytes,1,opt,name=new_password,json=newPassword" json:"new_password,omitempty"`
	OldPassword string `protobuf:"bytes,2,opt,name=old_password,json=oldPassword" json:"old_password,omitempty"`
	UserId      int64  `protobuf:"varint,3,opt,name=user_id,json=userId" json:"user_id,omitempty"`
}

func (m *ModifyPasswordRequest) Reset()                    { *m = ModifyPasswordRequest{} }
func (m *ModifyPasswordRequest) String() string            { return proto.CompactTextString(m) }
func (*ModifyPasswordRequest) ProtoMessage()               {}
func (*ModifyPasswordRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *ModifyPasswordRequest) GetNewPassword() string {
	if m != nil {
		return m.NewPassword
	}
	return ""
}

func (m *ModifyPasswordRequest) GetOldPassword() string {
	if m != nil {
		return m.OldPassword
	}
	return ""
}

func (m *ModifyPasswordRequest) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func init() {
	proto.RegisterType((*UserInfo)(nil), "light.services.user.v1.UserInfo")
	proto.RegisterType((*RegisterRequest)(nil), "light.services.user.v1.RegisterRequest")
	proto.RegisterType((*RegisterResponse)(nil), "light.services.user.v1.RegisterResponse")
	proto.RegisterType((*LoginRequest)(nil), "light.services.user.v1.LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "light.services.user.v1.LoginResponse")
	proto.RegisterType((*AuthenticateRequest)(nil), "light.services.user.v1.AuthenticateRequest")
	proto.RegisterType((*AuthenticateResponse)(nil), "light.services.user.v1.AuthenticateResponse")
	proto.RegisterType((*GetUserInfoRequest)(nil), "light.services.user.v1.GetUserInfoRequest")
	proto.RegisterType((*ModifyUserInfoRequest)(nil), "light.services.user.v1.ModifyUserInfoRequest")
	proto.RegisterType((*ModifyPasswordRequest)(nil), "light.services.user.v1.ModifyPasswordRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for UserService service

type UserServiceClient interface {
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	Authenticate(ctx context.Context, in *AuthenticateRequest, opts ...grpc.CallOption) (*AuthenticateResponse, error)
	GetUserInfo(ctx context.Context, in *GetUserInfoRequest, opts ...grpc.CallOption) (*UserInfo, error)
	ModifyUserInfo(ctx context.Context, in *ModifyUserInfoRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
	ModifyPassword(ctx context.Context, in *ModifyPasswordRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := grpc.Invoke(ctx, "/light.services.user.v1.UserService/Register", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := grpc.Invoke(ctx, "/light.services.user.v1.UserService/Login", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Authenticate(ctx context.Context, in *AuthenticateRequest, opts ...grpc.CallOption) (*AuthenticateResponse, error) {
	out := new(AuthenticateResponse)
	err := grpc.Invoke(ctx, "/light.services.user.v1.UserService/Authenticate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserInfo(ctx context.Context, in *GetUserInfoRequest, opts ...grpc.CallOption) (*UserInfo, error) {
	out := new(UserInfo)
	err := grpc.Invoke(ctx, "/light.services.user.v1.UserService/GetUserInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ModifyUserInfo(ctx context.Context, in *ModifyUserInfoRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/light.services.user.v1.UserService/ModifyUserInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ModifyPassword(ctx context.Context, in *ModifyPasswordRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/light.services.user.v1.UserService/ModifyPassword", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserService service

type UserServiceServer interface {
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	Authenticate(context.Context, *AuthenticateRequest) (*AuthenticateResponse, error)
	GetUserInfo(context.Context, *GetUserInfoRequest) (*UserInfo, error)
	ModifyUserInfo(context.Context, *ModifyUserInfoRequest) (*google_protobuf1.Empty, error)
	ModifyPassword(context.Context, *ModifyPasswordRequest) (*google_protobuf1.Empty, error)
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/light.services.user.v1.UserService/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/light.services.user.v1.UserService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Authenticate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthenticateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Authenticate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/light.services.user.v1.UserService/Authenticate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Authenticate(ctx, req.(*AuthenticateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/light.services.user.v1.UserService/GetUserInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserInfo(ctx, req.(*GetUserInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ModifyUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModifyUserInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ModifyUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/light.services.user.v1.UserService/ModifyUserInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ModifyUserInfo(ctx, req.(*ModifyUserInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ModifyPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModifyPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ModifyPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/light.services.user.v1.UserService/ModifyPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ModifyPassword(ctx, req.(*ModifyPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "light.services.user.v1.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _UserService_Register_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _UserService_Login_Handler,
		},
		{
			MethodName: "Authenticate",
			Handler:    _UserService_Authenticate_Handler,
		},
		{
			MethodName: "GetUserInfo",
			Handler:    _UserService_GetUserInfo_Handler,
		},
		{
			MethodName: "ModifyUserInfo",
			Handler:    _UserService_ModifyUserInfo_Handler,
		},
		{
			MethodName: "ModifyPassword",
			Handler:    _UserService_ModifyPassword_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}

func init() { proto.RegisterFile("user.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 597 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x55, 0xd1, 0x6e, 0xd3, 0x4a,
	0x10, 0x95, 0x9d, 0x3a, 0x4d, 0x26, 0xb9, 0xbd, 0x68, 0x09, 0x25, 0x72, 0x85, 0x08, 0x16, 0x88,
	0x88, 0x50, 0x47, 0x94, 0x07, 0xc4, 0x63, 0x91, 0x10, 0xaa, 0x04, 0x08, 0x85, 0x22, 0x24, 0x10,
	0x44, 0x6e, 0x3d, 0x49, 0x56, 0x38, 0x5e, 0xe3, 0xdd, 0x24, 0xf4, 0x2b, 0x78, 0xe5, 0x37, 0x78,
	0xe7, 0x47, 0xf8, 0x1b, 0xb4, 0x6b, 0xaf, 0x6b, 0xbb, 0x71, 0x03, 0x42, 0xe2, 0x71, 0x76, 0xce,
	0xcc, 0x9c, 0xb3, 0x7b, 0xc6, 0x06, 0x58, 0x70, 0x8c, 0xdd, 0x28, 0x66, 0x82, 0x91, 0xdd, 0x80,
	0x4e, 0x67, 0xc2, 0xe5, 0x18, 0x2f, 0xe9, 0x29, 0x72, 0x57, 0xa5, 0x96, 0x0f, 0xec, 0x9b, 0x53,
	0xc6, 0xa6, 0x01, 0x0e, 0x15, 0xea, 0x64, 0x31, 0x19, 0x0a, 0x3a, 0x47, 0x2e, 0xbc, 0x79, 0x94,
	0x14, 0xda, 0x7b, 0x65, 0x00, 0xce, 0x23, 0x71, 0x96, 0x24, 0x9d, 0x9f, 0x06, 0x34, 0xde, 0x70,
	0x8c, 0x8f, 0xc2, 0x09, 0x23, 0x3b, 0x60, 0x52, 0xbf, 0x6b, 0xf4, 0x8c, 0x7e, 0x6d, 0x64, 0x52,
	0x9f, 0xd8, 0xd0, 0x90, 0x53, 0x42, 0x6f, 0x8e, 0x5d, 0xb3, 0x67, 0xf4, 0x9b, 0xa3, 0x2c, 0x26,
	0x37, 0x00, 0x26, 0x34, 0xe6, 0x62, 0xac, 0xb2, 0x35, 0x95, 0x6d, 0xaa, 0x93, 0x97, 0x32, 0xbd,
	0x07, 0xcd, 0xc0, 0xd3, 0xd9, 0xad, 0xa4, 0x56, 0x1e, 0xa8, 0x64, 0x07, 0xac, 0x68, 0xc6, 0x42,
	0xec, 0x5a, 0x2a, 0x91, 0x04, 0xf2, 0x14, 0xe7, 0x1e, 0x0d, 0xba, 0xf5, 0xe4, 0x54, 0x05, 0xe4,
	0x31, 0x80, 0x6a, 0x14, 0xb0, 0x29, 0x0d, 0xbb, 0xdb, 0x3d, 0xa3, 0xdf, 0x3a, 0xb0, 0xdd, 0x44,
	0x92, 0xab, 0x25, 0xb9, 0xc7, 0x5a, 0xf3, 0x48, 0x8d, 0x7d, 0x2e, 0xc1, 0xce, 0x77, 0x03, 0xfe,
	0x1f, 0xe1, 0x94, 0x72, 0x81, 0xf1, 0x08, 0x3f, 0x2f, 0x90, 0x8b, 0x82, 0x24, 0xa3, 0x24, 0xc9,
	0x86, 0x46, 0xe4, 0x71, 0xbe, 0x62, 0xb1, 0xaf, 0xe5, 0xea, 0xf8, 0x5f, 0xc9, 0x75, 0x06, 0x70,
	0xe5, 0x9c, 0x32, 0x8f, 0x58, 0xc8, 0x91, 0x5c, 0x87, 0x6d, 0xc9, 0x71, 0x9c, 0xbd, 0x4d, 0x5d,
	0x86, 0x47, 0xbe, 0xf3, 0xd5, 0x80, 0xb6, 0x92, 0xfa, 0xb7, 0xea, 0xe4, 0xc3, 0x47, 0xa9, 0x2a,
	0x93, 0x46, 0xe4, 0x11, 0x34, 0xf1, 0x4b, 0x44, 0x63, 0x1c, 0x7b, 0x42, 0xc9, 0xb9, 0xfc, 0xce,
	0x1b, 0x09, 0xf8, 0x50, 0x38, 0x1f, 0xe1, 0xbf, 0x94, 0x50, 0xca, 0xbd, 0x03, 0x96, 0x60, 0x9f,
	0x30, 0x4c, 0xe9, 0x24, 0x41, 0xb1, 0xbf, 0xf9, 0x07, 0xfd, 0x07, 0x70, 0xf5, 0x70, 0x21, 0x66,
	0x18, 0x0a, 0x7a, 0xea, 0x09, 0xd4, 0xba, 0xd7, 0x4e, 0x71, 0x86, 0xd0, 0x29, 0x82, 0x37, 0xdd,
	0xe7, 0x3e, 0x90, 0x67, 0x28, 0xf4, 0x3a, 0xe8, 0xe6, 0x95, 0xf0, 0x6f, 0x06, 0x5c, 0x7b, 0xc1,
	0x7c, 0x3a, 0x39, 0x2b, 0x97, 0x14, 0xdd, 0x62, 0x5c, 0xea, 0x16, 0xb3, 0xca, 0x2d, 0xb5, 0xb5,
	0x6e, 0xd9, 0xca, 0x2f, 0x47, 0x8e, 0x9a, 0x55, 0xa0, 0xb6, 0xd4, 0xcc, 0x5e, 0xa5, 0x4f, 0xac,
	0x99, 0xdd, 0x82, 0x76, 0x88, 0xab, 0x71, 0xe6, 0x84, 0x84, 0x5b, 0x2b, 0xc4, 0x95, 0x46, 0x4a,
	0x08, 0x0b, 0xfc, 0x71, 0xc9, 0x2c, 0x2d, 0x16, 0xf8, 0x19, 0x24, 0x37, 0xb7, 0x96, 0x9f, 0x7b,
	0xf0, 0x63, 0x0b, 0x5a, 0xf2, 0x32, 0x5e, 0x27, 0x5f, 0x29, 0xf2, 0x01, 0x1a, 0xda, 0xce, 0xe4,
	0xae, 0xbb, 0xfe, 0x0b, 0xe6, 0x96, 0x76, 0xd4, 0xee, 0x6f, 0x06, 0xa6, 0x2f, 0x79, 0x0c, 0x96,
	0xb2, 0x1b, 0xb9, 0x5d, 0x55, 0x92, 0x5f, 0x0f, 0xfb, 0xce, 0x06, 0x54, 0xda, 0x95, 0x42, 0x3b,
	0xef, 0x1b, 0x32, 0xa8, 0x2a, 0x5b, 0x63, 0x45, 0xfb, 0xfe, 0xef, 0x81, 0xd3, 0x51, 0xef, 0xa1,
	0x95, 0x73, 0x1c, 0xb9, 0x57, 0x55, 0x7c, 0xd1, 0x96, 0x76, 0xaf, 0x0a, 0x9b, 0x75, 0x7b, 0x0b,
	0x3b, 0x45, 0x7b, 0x92, 0xfd, 0xaa, 0x9a, 0xb5, 0x36, 0xb6, 0x77, 0x2f, 0xec, 0xe4, 0x53, 0xf9,
	0xeb, 0x38, 0x6f, 0x9c, 0x19, 0x62, 0x43, 0xe3, 0x92, 0x0b, 0xab, 0x1a, 0x3f, 0xb1, 0xde, 0xd5,
	0xbc, 0x88, 0x9e, 0xd4, 0xd5, 0xf1, 0xc3, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xdd, 0x40, 0x5a,
	0xaf, 0xff, 0x06, 0x00, 0x00,
}
