BINARY_NAME=ragbag
BINARY_VERSION=next
SPEC_FILE_LOCATION=./resources/specs/ragbag-spec-v1.yml

build: clean generate
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
	rm -rf ./gen && mkdir -p ./gen/public && mkdir -p ./gen/restricted
	rm -rf ./frontend/src/gen && mkdir -p ./frontend/src/gen
	cd ./resources/db && sqlc generate && cd ../..
	docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli validate \
		-i "/local/${SPEC_FILE_LOCATION}"
	oapi-codegen -generate types,server -include-tags="auth" -package public ${SPEC_FILE_LOCATION} > ./gen/public/public.gen.go
	oapi-codegen -generate types,server -exclude-tags="auth" -package restricted ${SPEC_FILE_LOCATION} > ./gen/restricted/restricted.gen.go
	docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli generate \
		-i "/local/${SPEC_FILE_LOCATION}" \
		-g typescript-fetch \
		--additional-properties=typescriptThreePlus=true \
		-o "/local/frontend/src/gen"
