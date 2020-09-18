package rpc

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"io"
	"net"
	"time"
)

func ServeGRPC(grpcPort int, output io.Writer, route func(grpcServer *grpc.Server)) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.WithError(err).Fatalf("gRPC failed to listen: port=%d", grpcPort)
	}

	keepAliveOption := grpc.KeepaliveParams(keepalive.ServerParameters{
		Time: 10 * time.Second,
	})
	unaryInterceptorOption := grpc.ChainUnaryInterceptor(newLoggingInterceptor(output), errorAdapterInterceptor)
	gRPCServer := grpc.NewServer(keepAliveOption, unaryInterceptorOption)

	route(gRPCServer)
	if err := gRPCServer.Serve(listener); err != nil {
		log.WithError(err).Fatalln("gRPC 服务运行失败！")
	}
}
