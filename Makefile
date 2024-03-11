include .env
export

BINARY_NAME=todo

build: build-mac-arm #Replace for the default environment
.PHONY: build

build-linux:
	GOOS=linux GOARCH=amd64 go build -o bin/linux-amd64/${BINARY_NAME} ./cmd/api-server
.PHONY: build-linux

build-mac-amd:
	GOOS=darwin GOARCH=amd64 go build -o bin/darwin-amd64/${BINARY_NAME} ./cmd/api-server
.PHONY: build-mac-amd

build-mac-arm:
	GOOS=darwin GOARCH=arm64 go build -o bin/darwin-arm64/${BINARY_NAME} ./cmd/api-server
.PHONY: build-mac-arm

build-win:
	GOOS=windows GOARCH=amd64 go build -o bin/windows-amd64/${BINARY_NAME} ./cmd/api-server
.PHONY: build-win

clean:
	go clean
	go clean -cache
	rm -rf bin
.PHONY: clean

# generate: generate-openapi-files generate-di-files ### Generate all files

# generate-openapi-files: ### Generate server files based on OpenAPI specs
# 	openapi-generator generate -i docs/openapi.yaml -g go-gin-server  -o ./internal/interfaces/rest/v1 && rm -rf ./internal/interfaces/rest/v1/main.go ./internal/interfaces/rest/v1/go.mod ./internal/interfaces/rest/v1/Dockerfile ./internal/interfaces/rest/v1/go.sum ./internal/interfaces/rest/v1/api
# .PHONY: generate-openapi-files

# generate-di-files: ### Generate dependency injection files
# 	wire ./...
# .PHONY: generate-di-files

lint: ### check by golangci linter
	golangci-lint run
.PHONY: lint

migrate-up: ### migration up
	migrate -path migration -database '$(PG_URL)?sslmode=disable' up
.PHONY: migrate-up

run:
	go run ./cmd/api-server/.
.PHONY: run

setup-mac: ### setup mac os dependencies to run all tasks
	brew install golangci-lint
	brew install golang-migrate
	# brew install openapi-generator
	go install github.com/google/wire/cmd/wire@v0.6.0
.PHONY: setup-mac

test:
	go test -v -coverpkg=./internal/... ./tests/...
.PHONY: test

wire:
	wire ./internal/...
.PHONY: wire