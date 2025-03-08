BINARY_NAME=LaraApp

build:
	@echo "Building Lara..."
	@go build -o tmp/${BINARY_NAME} .
	@echo "Lara built!"

run: build
	@echo "Starting Lara..."
	@./tmp/${BINARY_NAME} &
	@echo "Lara started!"

clean:
	@echo "Cleaning..."
	@go clean
	@rm tmp/${BINARY_NAME}
	@echo "Cleaned!"

test:
	@echo "Testing..."
	@go test ./...
	@echo "Done!"

start: run

stop:
	@echo "Stopping Lara..."
	@-pkill -SIGTERM -f "./tmp/${BINARY_NAME}"
	@echo "Stopped Lara!"

restart: stop start
