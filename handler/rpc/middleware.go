package rpc

import (
	"context"
	"encoding/json"
	protoRPC "github.com/light-service/protobuf/rpc"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"io"
	"time"
)

func errorAdapterInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	resp, err = handler(ctx, req)
	grpcErr := protoRPC.GRPCError(err)
	return resp, grpcErr
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
