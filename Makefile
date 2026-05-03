# The name of the binary we want to build
BINARY_NAME=sage
# The entry point of our application
MAIN_PATH=main.go

# .PHONY tells Make that these aren't actual files, just command names
.PHONY: all build run clean test

# Default target when you just type 'make'
all: build

# Compiles the Cobra CLI into a single binary executable
build:
	@echo "Building $(BINARY_NAME)..."
	go build -o $(BINARY_NAME) $(MAIN_PATH)

# Builds the app and immediately runs it (great for testing)
run: build
	./$(BINARY_NAME) init

# Removes the compiled binary and any generated test files
clean:
	@echo "Cleaning up..."
	go clean
	rm -f $(BINARY_NAME)
	rm -f clonesage.yaml

# Runs all tests in the internal and cmd directories
test:
	@echo "Running tests..."
	go test ./... -v