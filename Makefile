# Variables
BINARY_NAME = delivery-optimiser
INPUT_FILE = input.json
VALIDATION_FILE = input_validation.json

# Build the binary
build:
	@go build -o $(BINARY_NAME) cmd/main.go

# Run the program (builds first, then executes)
run: build
	@./$(BINARY_NAME) $(INPUT_FILE) $(VALIDATION_FILE)
