# Chat cli Golang

## configuration 
* use ```CONFIGOR_ENV``` environment variable to specify which config to use. available configs are
1. config.dev.yml for testing and development environment, set **CONFIGOR_ENV=dev** (Default)
2. config.prod.yml for the production environment, set **CONFIGOR_ENV=prod**.

Please refer to config.template.yml file to see the available field for configuration.

## Prerequisites  
* Go v 1.14+
* protoc and protobuf
* protoc-gen-go-grpc module
* makefile

## Build and Run

## Locally

##### 1. Download project dependencies

    make init

##### 2. Generate chat grpc services

    make generate-proto

##### 3. Run tests
    
    make test
    
##### 4. Run test race
    
    make test-race

#### 5. build the server

    make build-server

    
#### 6. Run the server

    make run-server
    
#### 6. Run the client

    make run-client

## With Docker 

Build 

    make docker-build
Run 

    make docker-run

# what next ? 
* implement database layer to store users and chat history 
* create model for the user
* add grpc interceptors
* add grpc server health check
* add kubernetes deployment