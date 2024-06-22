# Makefile

# Variables
BINARY_NAME = timekeeper-backend
BUILD_DIR = bin

# Phony targets
.PHONY: run build clean test swag

# Run the application
run:
	@go run cmd/app/main.go

# Build the application
build:
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(BINARY_NAME) cmd/app/main.go

# Clean build artifacts
clean:
	@rm -rf $(BUILD_DIR)

# Run tests
test:
	@go test ./...

# Generate Swagger documentation
swag:
	@swag init -g cmd/app/main.go