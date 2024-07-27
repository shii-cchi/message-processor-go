include .env

.PHONY: all
all: build run

.PHONY: build
build:
	go build -o server cmd/main.go

.PHONY: run
run:
	./server

.PHONY: migration
migration:
	cd internal/database/schema && goose postgres ${DB_URI} up

.PHONY: migration_down
migration_down:
	cd internal/database/schema && goose postgres ${DB_URI} down

.PHONY: sqlc
sqlc:
	sqlc generate