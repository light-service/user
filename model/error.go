package model

import (
	"google.golang.org/grpc/codes"
)

const (
	ErrCodeOK      = 0
	ErrCodeUnknown = 99999
	errorCodeBase  = 11000
)

var (
	ErrNoAuth = newError(errorCodeBase+1, codes.Unauthenticated, "no auth")
	ErrNoPerm = newError(errorCodeBase+2, codes.PermissionDenied, "no perm")
)

type Error struct {
	errCode  int
	grpcCode codes.Code
	text     string
}

func newError(errCode int, grpcCode codes.Code, text string) Error {
	return Error{
		errCode:  errCode,
		grpcCode: grpcCode,
		text:     text,
	}
}

func (e Error) Code() int {
	return e.errCode
}

func (e Error) GRPCCode() codes.Code {
	return e.grpcCode
}

func (e Error) Error() string {
	return e.text
}

