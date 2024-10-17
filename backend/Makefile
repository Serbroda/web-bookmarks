BINARY_NAME=ragbag
BINARY_VERSION=next
ANGULAR_DIR=./ui/v1
SERVER_MAIN=./cmd/server/main.go

build: clean build-angular build-go-server

build-go-server:
	GOOS=darwin GOARCH=amd64 go build -ldflags "-X main.version=$(BINARY_VERSION)" -o build/${BINARY_NAME}-darwin-amd64 ${SERVER_MAIN}
	GOOS=linux GOARCH=amd64 go build -ldflags "-X main.version=$(BINARY_VERSION)" -o build/${BINARY_NAME}-linux-amd64 ${SERVER_MAIN}
	GOOS=windows GOARCH=amd64 go build -ldflags "-X main.version=$(BINARY_VERSION)" -o build/${BINARY_NAME}-windows-amd64.exe ${SERVER_MAIN}
	GOOS=linux GOARCH=arm GOARM=7 go build -ldflags "-X main.version=$(VERSION)" -o build/${BINARY_NAME}-linux-arm ${SERVER_MAIN}

generate-sqlc:
	find ./internal/db/sqlc -type f ! -name "mappings.go" -delete
	cd ./internal/db && sqlc generate && cd ../..

build-angular:
	cd "$(ANGULAR_DIR)" && npm install && npm run build

build-docker:
	docker build . -t ragbag:latest

clean:
	rm -rf "$(ANGULAR_DIR)/dist"
	rm -rf ./build
