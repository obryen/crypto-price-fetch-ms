APP_NAME=crypto-price-fetcher
BUILD_DIR=build
SRC_DIR=.

build:	clean
	go build -o $(BUILD_DIR)/$(APP_NAME) $(SRC_DIR)
run: build
	./$(BUILD_DIR)/$(APP_NAME)
test: 
	go test
proto:
	protoc --go_out=. --go_opt=paths=source_relative \
	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
	proto/service.proto

.PHONY: clean build run
	proto
