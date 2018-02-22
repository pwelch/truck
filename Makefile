default: proto

build: server client

server:
	go build server.go

client:
	go build -o cli client/client.go

clean:
	rm server cli output.txt

proto:
	protoc -I ./protobuf --go_out=plugins=grpc:./protobuf ./protobuf/*.proto

.PHONY: proto server client build clean
