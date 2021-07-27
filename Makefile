gen:
	protoc --go_out=. --go-grpc_out=. ./grpc_service/message.proto