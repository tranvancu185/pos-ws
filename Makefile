APP_NAME=server
GOOSE_DRIVER ?= sqlite3
GOOSE_DBSTRING="./veyPosDev.db"
GOOSE_MIGRATIONS_DIR ?= sql/schema

run:
	@go run ./cmd/$(APP_NAME)/server.go

build:
	@go build ./cmd/$(APP_NAME)/$(APP_NAME).go

sqlc:
	sqlc generate
	
upser:
	@GOOSE_DRIVER=sqlite3 GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir $(GOOSE_MIGRATIONS_DIR) up

downser:
	@GOOSE_DRIVER=sqlite3 GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir $(GOOSE_MIGRATIONS_DIR) down

resetser:
	@GOOSE_DRIVER=sqlite3 GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir $(GOOSE_MIGRATIONS_DIR) reset

.PHONY: upser downser resetser run build sqlc