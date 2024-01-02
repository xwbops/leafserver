# Makefile
GOPATH:=$(shell go env GOPATH)

# api 相关指令
.PHONY: api

api:
	# protoc --proto_path=./proto --go_out=. ./proto/msg.proto
	protoc --proto_path=./server/msg/proto --go_out=. ./server/msg/proto/msg.proto

