include .env

.DEFAULT_GOAL := build-and-run

BIN_FILE=main.out

PKGS := github.com/Soup666/diss-api,github.com/Soup666/diss-api/cmd/vision,github.com/Soup666/diss-api/controller,github.com/Soup666/diss-api/database,github.com/Soup666/diss-api/docs,github.com/Soup666/diss-api/middleware,github.com/Soup666/diss-api/middleware_test,github.com/Soup666/diss-api/model,github.com/Soup666/diss-api/repository,github.com/Soup666/diss-api/router,github.com/Soup666/diss-api/seed,github.com/Soup666/diss-api/seeder,github.com/Soup666/diss-api/seeds,github.com/Soup666/diss-api/services,github.com/Soup666/diss-api/utils

build-and-run: build run

build:
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o "${BIN_FILE}"

clean:
	go clean
	rm --force "cp.out"
	rm --force nohup.out

run:
	./"${BIN_FILE}"

test:
	@godotenv -f .env.test go test ./... -cover fmt

check:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 godotenv -f .env.test go test ./...

cover:
	@echo "Running tests with coverage..."
	@echo "${COVERPKG}"
	@godotenv -f .env.test go test -coverprofile=cp.out -coverpkg=$(PKGS) ./...
	@go tool cover -html=cp.out -o cover.html
	@go-cover-treemap -coverprofile=cp.out > out.svg

doc:
	swag init

lint:
	golangci-lint run --enable-all

run-seed:
	@GOOSE_DRIVER=postgres GOOSE_DBSTRING=$(DATABASE_URL) goose -dir=$(MIGRATION_PATH) reset
	@GOOSE_DRIVER=postgres GOOSE_DBSTRING=$(DATABASE_URL) goose -dir=$(MIGRATION_PATH) up
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go run seeder/seeder.go

db-status:
	@GOOSE_DRIVER=postgres GOOSE_DBSTRING=$(DATABASE_URL) goose -dir=$(MIGRATION_PATH) status

up:
	@GOOSE_DRIVER=postgres GOOSE_DBSTRING=$(DATABASE_URL) goose -dir=$(MIGRATION_PATH) up

reset:
	@GOOSE_DRIVER=postgres GOOSE_DBSTRING=$(DATABASE_URL) goose -dir=$(MIGRATION_PATH) reset

run-test:
	- make test
	- make open-test