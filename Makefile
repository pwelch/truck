default: proto

build:
	go build server.go

proto:
	protoc -I ./protobuf --go_out=plugins=grpc:./protobuf ./protobuf/*.proto

.PHONY: proto
