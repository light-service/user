protos:
	cd api/v1 && protoc --go_out=plugins=grpc:. *.proto

mocks:
	GO111MODULE=on mockgen -package=mocks -destination=api/v1/client/mocks/user_service_client_mock.go github.com/light-service/user/api/v1/client UserServiceClient
