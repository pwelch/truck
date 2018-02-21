default: proto

proto:
	protoc -I ./protobuf --go_out=plugins=grpc:./protobuf ./protobuf/*.proto

.PHONY: proto
