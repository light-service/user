package rpc

import (
	"context"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"time"
)

func adapterError(err error) error {
	if err == nil {
		return nil
	}

	grpcCode := codes.Internal // 未知错误
	if err, ok := err.(interface {
		GRPCCode() codes.Code
	}); ok {
		grpcCode = err.GRPCCode()
	}

	return status.New(grpcCode, err.Error()).Err()
}

func errorAdapterInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	resp, err = handler(ctx, req)
	err = adapterError(err)
	return resp, err
}

func newLoggingInterceptor(out io.Writer) grpc.UnaryServerInterceptor {
	log := logrus.New()
	log.SetOutput(out)
	log.SetFormatter(&logrus.JSONFormatter{TimestampFormat: time.RFC3339Nano})

	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		since := time.Now()
		defer func() {
			elapsed := time.Since(since)
			reqJson, _ := json.Marshal(req)
			respJson, _ := json.Marshal(resp)
			log.Infof("[grpc access] %s: req=%s, resp=%s, elapsed=%.2f", info.FullMethod, reqJson, respJson, float64(elapsed)/float64(time.Second))
		}()

		resp, err = handler(ctx, req)
		return resp, err
	}
}
