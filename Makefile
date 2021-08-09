VERSION=1.0.0

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GORUN=$(GOCMD) run
GOMOD=$(GOCMD) mod
GOTOOL=$(GOCMD) tool
DEFAULT_DB=mongodb

# Server name and path
BINARY_SERVER_NAME=chat-cli-go
SERVER_PATH=./src/cmd/server/server.go
CLIENT_PATH=./src/cmd/client/client.go


init:
# todo : upgrade grpc-gateway lib
	$(GOGET) -u github.com/grpc-ecosystem/grpc-gateway
	$(GOGET) -u google.golang.org/grpc
	$(GOGET) -u github.com/golang/protobuf/protoc-gen-go
	$(GOGET) -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	$(GOGET) -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
	$(GOMOD) download

# TESTS AND COVERAGE
test:
	$(GOTEST)  -v ./...

test-race:
	$(GOTEST) --race -v ./...

coverage:
	$(GOTEST)  ./... -coverprofile cover.out
	go tool cover -html=cover.out

coverage-badge:
	gopherbadger -md="README.md"

# BUILD
build-server:
	$(GOBUILD) -o ./bin/$(BINARY_SERVER_NAME) -v $(SERVER_PATH)

build: test build-server

# RUN
run-server:
	$(GORUN)  $(SERVER_PATH)

run-client:
	$(GORUN)  $(CLIENT_PATH)

docker-build:
	docker build \
	-t gharsallahmoez/chat-cli-go:${VERSION} .

docker-run:
	docker run gharsallahmoez/chat-cli-go:${VERSION}

vet:
	$(GOCMD) vet ./...

# CLEAN
.PHONY: clean
clean:
	$(GOCLEAN)
	$(GOCLEAN) -testcache

clean-proto:
	rm -rf ./src/pb

generate-proto:
	protoc -I src/protobuf --go_out=src/pb --go-grpc_out=src/pb --go-grpc_opt=paths=source_relative chat.proto


