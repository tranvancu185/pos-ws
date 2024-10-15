-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS tables (
  "table_id" INTEGER PRIMARY KEY AUTOINCREMENT,
  "table_name" VARCHAR(255) NOT NULL,
  "table_code" VARCHAR(5) NOT NULL UNIQUE,
  "table_status" INTEGER NOT NULL DEFAULT 1,
  "table_properties" TEXT DEFAULT NULL,
  "created_at" INTEGER DEFAULT 0,
  "updated_at" INTEGER DEFAULT 0,
  "deleted_at" INTEGER DEFAULT 0
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS tables;
-- +goose StatementEnd
