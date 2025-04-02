# Variables
APP_NAME := nats-client
BUILD_DIR := build
SRC_DIR := ./cmd

# Default target
.PHONY: all
all: build

# Build command
.PHONY: build
build:
	@echo "Building the application..."
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(APP_NAME) $(SRC_DIR)

# Clean command
.PHONY: clean
clean:
	@echo "Cleaning up build artifacts..."
	@rm -rf $(BUILD_DIR)

# Run command
.PHONY: run
run:
	@echo "Running the application..."
	@go run $(SRC_DIR)/...

# Test command
.PHONY: test
test:
	@echo "Running tests..."
	@go test ./... -v
