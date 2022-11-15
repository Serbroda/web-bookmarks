BINARY_NAME=ragbag
BINARY_VERSION=next
SPEC_FILE_LOCATION=./resources/specs/ragbag-spec-v1.yml

build: clean
	cd frontend && yarn install && yarn build && cd ..
	GOOS=darwin GOARCH=amd64 go build -ldflags "-X main.version=$(BINARY_VERSION)" -o build/${BINARY_NAME}-darwin-amd64 main.go
	GOOS=linux GOARCH=amd64 go build -ldflags "-X main.version=$(BINARY_VERSION)" -o build/${BINARY_NAME}-linux-amd64 main.go
	GOOS=windows GOARCH=amd64 go build -ldflags "-X main.version=$(BINARY_VERSION)" -o build/${BINARY_NAME}-windows-amd64.exe main.go
#	GOOS=linux GOARCH=arm GOARM=7 go build -ldflags "-X main.version=$(VERSION)" -o build/${BINARY_NAME}-linux-arm main.go

clean:
	go clean
	rm -rf frontend/dist/
	rm -rf build/

generate:
	docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli validate \
		-i "/local/${SPEC_FILE_LOCATION}"
	rm -rf ./gen && mkdir ./gen && oapi-codegen -generate types,server -package gen ${SPEC_FILE_LOCATION} > ./gen/ragbag.gen.go
	rm -rf ./frontend/src/gen && mkdir ./frontend/src/gen && docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli generate \
		-i "/local/${SPEC_FILE_LOCATION}" \
		-g typescript-fetch \
		--additional-properties=typescriptThreePlus=true \
		-o "/local/frontend/src/gen"
