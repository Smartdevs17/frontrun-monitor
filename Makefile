.PHONY: help build run test clean fmt vet lint deps install dev

# Variables
BINARY_NAME=frontrun-monitor
BACKEND_DIR=backend
BINARY_PATH=$(BACKEND_DIR)/$(BINARY_NAME)
GO_VERSION=1.21

# Default target
help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Available targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

# Build targets
build: ## Build the backend binary
	@echo "Building $(BINARY_NAME)..."
	@cd $(BACKEND_DIR) && go build -o $(BINARY_NAME) .
	@echo "Build complete: $(BINARY_PATH)"

build-linux: ## Build the backend binary for Linux
	@echo "Building $(BINARY_NAME) for Linux..."
	@cd $(BACKEND_DIR) && GOOS=linux GOARCH=amd64 go build -o $(BINARY_NAME)-linux .
	@echo "Build complete: $(BACKEND_DIR)/$(BINARY_NAME)-linux"

build-windows: ## Build the backend binary for Windows
	@echo "Building $(BINARY_NAME) for Windows..."
	@cd $(BACKEND_DIR) && GOOS=windows GOARCH=amd64 go build -o $(BINARY_NAME).exe .
	@echo "Build complete: $(BACKEND_DIR)/$(BINARY_NAME).exe"

build-darwin: ## Build the backend binary for macOS
	@echo "Building $(BINARY_NAME) for macOS..."
	@cd $(BACKEND_DIR) && GOOS=darwin GOARCH=amd64 go build -o $(BINARY_NAME)-darwin .
	@echo "Build complete: $(BACKEND_DIR)/$(BINARY_NAME)-darwin"

build-all: build-linux build-windows build-darwin ## Build for all platforms

# Run targets
run: ## Run the backend server
	@echo "Running $(BINARY_NAME)..."
	@cd $(BACKEND_DIR) && go run .

dev: ## Run the backend server in development mode with auto-reload (requires air)
	@which air > /dev/null || (echo "Error: 'air' not found. Install it with: go install github.com/cosmtrek/air@latest" && exit 1)
	@cd $(BACKEND_DIR) && air

# Development tools
deps: ## Download and install dependencies
	@echo "Installing core Ethereum libraries..."
	@cd $(BACKEND_DIR) && go get github.com/ethereum/go-ethereum
	@echo "Installing additional libraries..."
	@cd $(BACKEND_DIR) && go get github.com/gorilla/websocket
	@cd $(BACKEND_DIR) && go get github.com/sirupsen/logrus
	@cd $(BACKEND_DIR) && go get golang.org/x/time/rate
	@echo "Downloading dependencies..."
	@cd $(BACKEND_DIR) && go mod download
	@cd $(BACKEND_DIR) && go mod tidy
	@echo "Dependencies installed"

fmt: ## Format Go code
	@echo "Formatting code..."
	@cd $(BACKEND_DIR) && go fmt ./...
	@echo "Formatting complete"

vet: ## Run go vet
	@echo "Running go vet..."
	@cd $(BACKEND_DIR) && go vet ./...
	@echo "Vet complete"

lint: ## Run golangci-lint (if installed)
	@which golangci-lint > /dev/null || (echo "Warning: 'golangci-lint' not found. Skipping lint." && exit 0)
	@echo "Running golangci-lint..."
	@cd $(BACKEND_DIR) && golangci-lint run
	@echo "Lint complete"

test: ## Run tests
	@echo "Running tests..."
	@cd $(BACKEND_DIR) && go test -v ./...
	@echo "Tests complete"

test-coverage: ## Run tests with coverage
	@echo "Running tests with coverage..."
	@cd $(BACKEND_DIR) && go test -v -coverprofile=coverage.out ./...
	@cd $(BACKEND_DIR) && go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: $(BACKEND_DIR)/coverage.html"

# Cleanup targets
clean: ## Clean build artifacts
	@echo "Cleaning build artifacts..."
	@rm -f $(BACKEND_DIR)/$(BINARY_NAME)
	@rm -f $(BACKEND_DIR)/$(BINARY_NAME)-linux
	@rm -f $(BACKEND_DIR)/$(BINARY_NAME)-darwin
	@rm -f $(BACKEND_DIR)/$(BINARY_NAME).exe
	@rm -f $(BACKEND_DIR)/coverage.out
	@rm -f $(BACKEND_DIR)/coverage.html
	@echo "Clean complete"

# Install target
install: build ## Build and install the binary
	@echo "Installing $(BINARY_NAME)..."
	@cp $(BINARY_PATH) /usr/local/bin/$(BINARY_NAME) || echo "Note: Installation to /usr/local/bin requires sudo. Run: sudo make install"
	@echo "Installation complete"

# Docker targets (optional)
docker-build: ## Build Docker image
	@echo "Building Docker image..."
	@docker build -t $(BINARY_NAME):latest -f Dockerfile .
	@echo "Docker image built"

docker-run: ## Run Docker container
	@echo "Running Docker container..."
	@docker run -p 8080:8080 --env-file .env $(BINARY_NAME):latest

# Check Go version
check-go: ## Check if Go version matches requirements
	@echo "Checking Go version..."
	@go version | grep -q "go$(GO_VERSION)" || (echo "Warning: Go version mismatch. Required: $(GO_VERSION)" && exit 1)
	@echo "Go version check passed"

