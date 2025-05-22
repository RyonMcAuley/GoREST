BINARY_NAME="myapp"

MAIN_DIR=./cmd/server/main.go

run:
	@echo "Running the application..."
	@go run $(MAIN_DIR)

build:
	@echo "Building the application..."
	@go build -o $(BINARY_NAME) $(MAIN_DIR)

clean:
	@echo "Cleaning up..."
	@rm -f $(BINARY_NAME)
