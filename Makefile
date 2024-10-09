# Go parameters
GOCMD = go
GOBUILD = $(GOCMD) build
GOTEST = $(GOCMD) test
GOFMT = $(GOCMD) fmt
BINARY_NAME = fc-codes

# Directories
CMD_DIR = ./cmd
INTERNAL_DIR = ./internal
BIN_DIR = ./bin

# Build the project
all: test build

# Format the Go source files
fmt:
	$(GOFMT) ./...

# Build the binary
build:
	$(GOBUILD) -o $(BIN_DIR)/$(BINARY_NAME) $(CMD_DIR)/main.go

# Run the app
run:
	$(GOCMD) run $(CMD_DIR)/main.go

# Run tests
test:
	$(GOTEST) ./...

# Clean the build
clean:
	rm -f $(BIN_DIR)/$(BINARY_NAME)

# Help command
help:
	@echo "Usage:"
	@echo "  make all          - Run tests and build"
	@echo "  make fmt          - Format source files"
	@echo "  make build        - Build binary"
	@echo "  make run          - Run the application"
	@echo "  make test         - Run tests"
	@echo "  make clean        - Clean build artifacts"
	@echo "  make build-linux  - Cross-compile for Linux"

.PHONY: all fmt build run test clean build-linux help
