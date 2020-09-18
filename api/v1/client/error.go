package client

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Code(err error) codes.Code {
	return status.Code(err)
}
