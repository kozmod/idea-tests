BIN_DIR=bin
SERVER_BIN=server
CLIENT_BIN=client

SERVER_SRC=simple-server/cmd
CLIENT_SRC=simple-client/cmd

env := GO111MODULE="on"

all: client server

protoc:
	@echo "Generating Go files"
	mkdir -p ./proto/generated
	protoc -I ./proto/file --go_out=plugins=grpc:./proto/generated api/api.proto

server: protoc
	@echo "Building server"
	mkdir -p ${BIN_DIR}
	pushd ${SERVER_SRC} && go build -o ../../${BIN_DIR}/${SERVER_BIN}

client: protoc
	@echo "Building client"
	mkdir -p ${BIN_DIR}
	pushd ${CLIENT_SRC} && go build -o ../../${BIN_DIR}/${CLIENT_BIN}

run.server:
	 pushd ${SERVER_SRC} && ${env} go run ${SERVER_BIN}.go
run.client:
	 pushd ${CLIENT_SRC} && ${env} go run ${CLIENT_BIN}.go

run.bin.server:
	pushd ${BIN_DIR} && ./${SERVER_BIN}
run.bin.client:
	pushd ${BIN_DIR} && ./${CLIENT_BIN}

.PHONY: client server protoc

clear:
	@(rm -rf ${BIN_DIR})

download:
	go mod download

imports:
	goimports -w .
