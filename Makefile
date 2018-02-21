default: proto

server:
	go build server.go

client:
	go build -o foo client/client.go

proto:
	protoc -I ./protobuf --go_out=plugins=grpc:./protobuf ./protobuf/*.proto

.PHONY: proto server client
