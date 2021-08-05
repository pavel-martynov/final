gen:
	protoc --go_out=. --go-grpc_out=. ./grpc_service/message.proto

# you can add here call for a linters and tests