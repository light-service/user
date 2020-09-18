package main

import (
	"github.com/light-service/user/handler/rpc"
	"github.com/light-service/user/utils"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"io"
	"os"
	"path/filepath"
	"time"
)

var (
	// 基本配置
	mysqlHost   = utils.Env("MYSQL_HOST", "localhost")
	mysqlPort   = utils.EnvInt("MYSQL_PORT", 3306)
	mysqlUser   = utils.Env("MYSQL_USER", "root")
	mysqlPass   = utils.Env("MYSQL_PASS", "")
	mysqlDBName = utils.Env("MYSQL_DBNAME", "user")
	grpcPort    = utils.EnvInt("GRPC_PORT", 8001)
	dataPath    = utils.Env("DATA_PATH", "/data/")
	logLevel    = utils.Env("LOG_LEVEL", "info")
)

func main() {
	mainLogFile := openFile(filepath.Join(dataPath, "main.log"))
	defer mainLogFile.Close()
	grpcLogFile := openFile(filepath.Join(dataPath, "grpc.log"))
	defer grpcLogFile.Close()

	configLog(mainLogFile, logLevel)
	rpc.ServeGRPC(grpcPort, grpcLogFile, func(grpcServer *grpc.Server) {
		rpc.Serve(grpcServer)
	})
}

func openFile(filePath string) *os.File {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.WithError(err).Fatalf("failed to open file: %s", filePath)
	}
	return file
}

func configLog(out io.Writer, level string) {
	if l, err := log.ParseLevel(level); err == nil {
		log.SetLevel(l)
	} else {
		log.WithError(err).Errorf("parse level failed: level=%s", level)
	}
	log.SetOutput(out)
	log.SetFormatter(&log.JSONFormatter{TimestampFormat: time.RFC3339Nano})
	log.SetReportCaller(true)
}
