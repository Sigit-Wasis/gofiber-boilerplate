.PHONY: build run migrate-up migrate-down seed swag-init deps clean fmt

APP_NAME=app

# Load .env if exists
ifneq (,$(wildcard ./.env))
    include .env
    export $(shell sed 's/=.*//' .env)
endif

build:
	go build -o $(APP_NAME) ./

run: swag-init
	go run main.go

migrate-up:
	migrate -path ./migrations -database "$(DATABASE_URL)" up

migrate-down:
	migrate -path ./migrations -database "$(DATABASE_URL)" down

reset-db:
	psql -d postgres -c "SELECT pg_terminate_backend(pg_stat_activity.pid) \
	FROM pg_stat_activity \
	WHERE pg_stat_activity.datname = 'gofiber-boilerplate' AND pid <> pg_backend_pid();"
	dropdb --if-exists gofiber-boilerplate
	createdb gofiber-boilerplate

swag-init:
	swag init -g main.go -o docs

deps:
	go mod tidy
	@echo "Install swag: go install github.com/swaggo/swag/cmd/swag@latest"
	@echo "Install migrate: go install github.com/golang-migrate/migrate/v4/cmd/migrate@latest"

clean:
	rm -f $(APP_NAME)
	rm -rf docs

fmt:
	gofmt -s -w .
