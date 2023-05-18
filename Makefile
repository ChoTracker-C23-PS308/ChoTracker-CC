include .env.example
-include .env

BIN=web
epoch=$(shell date +%s)

.PHONY: all
all: init test clean-bin build

.PHONY: setup
setup: env init

.PHONY: init
init:
	@echo ">> Initializing dependencies"
	@go mod tidy -v
	@go get -v ./...

.PHONY: test
test:
	@echo ">> Testing all"
	@go test -cover -covermode atomic -coverprofile cover.out -race ./...
	@go tool cover -func cover.out

.PHONY: build
build:
	@echo ">> Building ${BIN} binary"
	@go build -o bin/${BIN} ./cmd/${BIN}

.PHONY: run
run:
	@echo ">> Running ${BIN} binary"
	@bin/${BIN}

run3:
	@bash -c 'set -o allexport; source .env; set +o allexport; go run cmd/web/main.go'

.PHONY: clean-bin
clean-bin:
	@echo ">> Removing binary directory"
	@rm -rf bin

.PHONY: env
env:
	@echo ">> Creating .env"
	@cat .env.example > .env
	@echo ">> Creating .env.local"
	@cat .env.example > .env.local

.PHONY: schema
schema:
	@touch $(DATABASE_MIGRATIONS_PATH)"/"$(epoch)"_[RENAME].up.sql"
	@touch $(DATABASE_MIGRATIONS_PATH)"/"$(epoch)"_[RENAME].down.sql"
	@echo ">> Created schema in "$(DATABASE_MIGRATIONS_PATH)
	@echo "\t- "$(epoch)"_[RENAME].up.sql"
	@echo "\t- "$(epoch)"_[RENAME].down.sql"
	@echo ">> Please rename the generated files"

.PHONY: web
web:
	@docker-compose up --build -d web

.PHONY: db
db:
	@docker-compose up --build -d db


.PHONY: migrate-up
migrate-up:
	@migrate -source file://${DATABASE_MIGRATIONS_PATH} -database ${DATABASE_URL} up

.PHONY: migrate-down
migrate-down:
	@migrate -source file://${DATABASE_MIGRATIONS_PATH} -database ${DATABASE_URL} down 1

.PHONY: migrate
migrate:
	@migrate -source file://${DATABASE_MIGRATIONS_PATH} -database ${DATABASE_URL} version