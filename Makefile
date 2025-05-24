BINARY_NAME="myapp"

MAIN_DIR=./cmd/server/main.go

run:
	@echo "Running the application..."
	@go run $(MAIN_DIR)

build:
	@echo "Building the application..."
	@go build -o $(BINARY_NAME) $(MAIN_DIR)

docker-build:
	@echo "Building Docker image..."
	@docker build -t $(BINARY_NAME) .

docker-run:
	@echo "Running Docker container..."
	@docker run -d --name $(BINARY_NAME) -p 8080:8080 $(BINARY_NAME)

docker-restart:
	@echo "Restarting Docker container..."
	@docker restart $(BINARY_NAME)

docker-stop:
	@echo "Stopping Docker container..."
	@docker stop $(BINARY_NAME)

start:
	@echo "Starting the application..."
	@./$(BINARY_NAME)

clean:
	@echo "Cleaning up..."
	@rm -f $(BINARY_NAME)
