BINARY_NAME=ragbag
BINARY_VERSION=next
SPEC_FILE_LOCATION=./app/resources/specs/ragbag-spec-v1.yml
SWAGGER_SERVER_TARGET_DIR=./app/cmd/rest-server/handlers

build: clean generate
	cd frontend && yarn install && yarn build && cd ..
	GOOS=darwin GOARCH=amd64 go build -ldflags "-X main.version=$(BINARY_VERSION)" -o build/${BINARY_NAME}-darwin-amd64 app/cmd/rest-server/main.go
	GOOS=linux GOARCH=amd64 go build -ldflags "-X main.version=$(BINARY_VERSION)" -o build/${BINARY_NAME}-linux-amd64 app/cmd/rest-server/main.go
	GOOS=windows GOARCH=amd64 go build -ldflags "-X main.version=$(BINARY_VERSION)" -o build/${BINARY_NAME}-windows-amd64.exe app/cmd/rest-server/main.go
#	GOOS=linux GOARCH=arm GOARM=7 go build -ldflags "-X main.version=$(VERSION)" -o build/${BINARY_NAME}-linux-arm main.go

clean:
	go clean
	rm -rf frontend/dist/
	rm -rf build/

generate: generate-sqlc generate-swagger-server generate-swagger-client


generate-sqlc:
	rm -rf ./app/sqlc
	cd ./app/resources/db && sqlc generate && cd ../..

generate-swagger-server: validate-swagger
	mkdir -p ${SWAGGER_SERVER_TARGET_DIR}/public
	mkdir -p ${SWAGGER_SERVER_TARGET_DIR}/restricted
	oapi-codegen -generate types,server -include-tags="auth" -package public \
		${SPEC_FILE_LOCATION} > ${SWAGGER_SERVER_TARGET_DIR}/public/gen.go
	oapi-codegen -generate types,server -exclude-tags="auth" -package restricted \
		${SPEC_FILE_LOCATION} > ${SWAGGER_SERVER_TARGET_DIR}/restricted/gen.go

generate-swagger-client: validate-swagger
	docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli generate \
		-i "/local/${SPEC_FILE_LOCATION}" \
		-g typescript-fetch \
		--additional-properties=typescriptThreePlus=true \
		-o "/local/frontend/src/gen"

validate-swagger:
	docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli validate \
		-i "/local/${SPEC_FILE_LOCATION}"
