.PHONY: build run migrate-up migrate-down seed swag-init deps clean fmt

APP_NAME=app

build:
	 go build -o $(APP_NAME) ./

run:
	 go run main.go

migrate-up:
	 migrate -path ./migrations -database "$${DATABASE_URL}" up

migrate-down:
	 migrate -path ./migrations -database "$${DATABASE_URL}" down

seed:
	 go run ./cmd/seed

swag-init:
	 swag init -g main.go -o docs

deps:
	 go mod tidy
	 @echo "Install swag: go install github.com/swaggo/swag/cmd/swag@latest"
	 @echo "Install migrate: brew install golang-migrate" || true

clean:
	 rm -f $(APP_NAME)
	 rm -rf docs

fmt:
	 gofmt -s -w .
