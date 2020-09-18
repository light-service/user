package client

import (
	"github.com/light-service/protobuf/rpc"
	"google.golang.org/grpc/codes"
)

func GRPCStatus(err error) codes.Code {
	return rpc.GRPCStatus(err)
}

func GRPCErrCode(err error) int {
	return rpc.GRPCErrCode(err)
}
