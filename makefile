APP_NAME=crypto-price-fetcher
BUILD_DIR=build
SRC_DIR=.

.PHONY: clean build run
build:	clean
	go build -o $(BUILD_DIR)/$(APP_NAME) $(SRC_DIR)
run: build
	./$(BUILD_DIR)/$(APP_NAME)
test: 
	go test
