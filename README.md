## =======> To setup server <=========

# Run command to setup server

go mod tidy
go mod verify
go mod vendor

# To run the server

go run cmd/server/server.go

# To build the server

go build cmd/server/server.go

# ====> Output build in root server

# cd to "server/tests/basic" and run this command to generate testing coverage ( do bao phu code )

go test -coverprofile=coverage

# parse to html for UI

go tool cover -html=coverage -o coverage.html

## ==========> MIGRATION DB <===========

# Install GOOSE ==> https://pkg.go.dev/github.com/pressly/goose/v3#section-readme

go install github.com/pressly/goose/v3/cmd/goose@latest

# To create a migration file

goose create <tên_migration> sql

<!-- goose create table_user sql -->

# To create a migration file with path

goose -dir <path> create table_user sql

<!-- goose -dir db/sql/schema create table_user sql -->

# To run migrations file

goose -dir <path_folder_migrations> sqlite3 <path_db_file>

<!-- goose -dir migrations sqlite3 /path/to/your/database.db up -->

## =========> SET UP ROUTER <==========

# Install WIRE ==> https://github.com/google/wire

go install github.com/google/wire/cmd/wire@latest

# cd to internal/wire and run this command

wire
