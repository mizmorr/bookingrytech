.ONESHELL:

SHELL := /bin/bash

POSTGRES_DB				?= books
POSTGRES_USER			?= postgres
POSTGRES_PASSWORD ?= post

HTTP_HOST					?= localhost
HTTP_PORT					?= 8080
GREEN=\033[0;32m
RESET=\033[0m

start-with-postgres:
	@trap 'make rollback-migrations' INT TERM
	@make migrate-up
	@echo "Service is starting!"
	@sleep 1
	@echo -e "$(GREEN)Open http://$(HTTP_HOST):$(HTTP_PORT)/swagger/ in browser or run make swagger-open in other terminal!$(RESET)"
	@make run

run:
	cd cmd; go run main.go

test:
	go test -v -cover ./...

swag:
	swag init -g cmd/main.go -o docs

migrate-up:
	@echo "migrations are going up.."
	migrate -path internal/store/postgres/migrations -database "postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@localhost:5432/$(POSTGRES_DB)?sslmode=disable" up

open-swagger:
	@xdg-open http://$(HTTP_HOST):$(HTTP_PORT)/swagger/

rollback-migrations:
	@echo "migrations are going down.."
	migrate -path internal/store/postgres/migrations -database "postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@localhost:5432/$(POSTGRES_DB)?sslmode=disable" down

