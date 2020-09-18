protos:
	# cd api/v1 && protoc --go-grpc_out=. *.proto
	cd api/v1 && protoc \
      --go_out=Mgrpc/service_config/service_config.proto=/internal/proto/grpc_service_config:. \
      --go-grpc_out=Mgrpc/service_config/service_config.proto=/internal/proto/grpc_service_config:. \
      --go_opt=paths=source_relative \
      --go-grpc_opt=paths=source_relative \
      user.proto

mocks:
	GO111MODULE=on mockgen -package=mocks -destination=api/v1/client/mocks/user_service_client_mock.go github.com/light-service/user/api/v1/client UserServiceClient
