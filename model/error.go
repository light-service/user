package model

import (
	"github.com/light-service/protobuf/rpc"
	"google.golang.org/grpc/codes"
)

const errorCodeBase = 11000

var (
	ErrNoAuth = rpc.NewError(errorCodeBase+1, codes.Unauthenticated, "no auth")
	ErrNoPerm = rpc.NewError(errorCodeBase+2, codes.PermissionDenied, "no perm")
)
