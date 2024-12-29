SHELL := /bin/bash

# ------------------------------------------------------------
# Projektkonfiguration
# ------------------------------------------------------------
PROJECT_NAME     := my-bookmark-app
BINARY_NAME      := server

# Pfade und Dateien
OPENAPI_FILE     := openapi.yml
GEN_OUTPUT_DIR   := internal/api

# Docker-Image für OpenAPI Generator
OPENAPI_DOCKER_IMAGE := openapitools/openapi-generator-cli

# ------------------------------------------------------------
# Standard-Targets
# ------------------------------------------------------------
.PHONY: all clean build test run generate docker-build docker-run docker-stop help

## Default: baut das Projekt
all: build

## Entfernt erzeugte Dateien
clean:
	@echo "==> Cleaning up..."
	rm -rf bin/$(BINARY_NAME)
	rm -rf $(GEN_OUTPUT_DIR)

## Erzeugt Code aus OpenAPI-Spezifikation
# https://github.com/oapi-codegen/oapi-codegen
generate:
	@echo "==> Generating Go code via oapi-codegen..."
	mkdir -p $(GEN_OUTPUT_DIR)
	go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest \
 		--config=oapi-codegen.yml \
 		$(OPENAPI_FILE)
	@echo "==> OpenAPI generation done. Check '$(GEN_OUTPUT_DIR)' folder."

## Baut das Go-Binary
build:
	@echo "==> Building Go binary..."
	go build -o bin/$(BINARY_NAME) ./cmd/server
	@echo "==> Build complete! Binary in bin/$(BINARY_NAME)."

## Führt alle Tests aus
test:
	@echo "==> Running tests..."
	go test ./... -v

## Startet den Server lokal (ohne Docker)
run: build
	@echo "==> Starting local server..."
	./bin/$(BINARY_NAME)

# ------------------------------------------------------------
# Docker-Targets
# ------------------------------------------------------------

## Baut das Docker-Image
docker-build:
	@echo "==> Building Docker image..."
	docker build -t $(PROJECT_NAME):latest .

## Startet Docker-Container (App + DB via docker-compose)
docker-run:
	@echo "==> Starting docker-compose..."
	docker-compose up --build

## Stoppt docker-compose
docker-stop:
	@echo "==> Stopping docker-compose..."
	docker-compose down

# ------------------------------------------------------------
# Hilfetext
# ------------------------------------------------------------
help:
	@echo "----------------------------------------------------------------"
	@echo "Available make targets:"
	@echo "  make all            -> Baut das Projekt (alias für 'build')"
	@echo "  make clean          -> Entfernt erzeugte Dateien"
	@echo "  make generate       -> Erzeugt Go-Code aus OpenAPI-Datei (via Docker)"
	@echo "  make build          -> Baut das Go-Binary"
	@echo "  make test           -> Führt Tests aus"
	@echo "  make run            -> Startet den Server lokal (ohne Docker)"
	@echo "  make docker-build   -> Baut das Docker-Image"
	@echo "  make docker-run     -> Startet das Projekt via docker-compose"
	@echo "  make docker-stop    -> Stoppt docker-compose und Container"
	@echo "  make help           -> Zeigt diese Hilfe"
	@echo "----------------------------------------------------------------"