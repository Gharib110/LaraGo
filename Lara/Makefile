## test: runs all tests
test:
	@go test -v ./...

## cover: opens coverage in browser
cover:
	@go test -coverprofile=coverage.out ./... && go tool cover -html=coverage.out

## coverage: displays test coverage
coverage:
	@go test -cover ./...

## build_cli: builds the command line tool Lara and copies it to test-app
build_cli:
	@go build -o ../test-app/lara ./cmd/cli

## build: builds the command line tool dist directory
build:
	@go build -o ./dist/lara ./cmd/cli

install_cli:
	@go build -o ~/go/bin/lara -ldflags '-s -w' ./cmd/cli