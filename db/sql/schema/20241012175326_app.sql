-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS apps (
  "app_id" INTEGER PRIMARY KEY AUTOINCREMENT,
  "app_name" VARCHAR NOT NULL,
  "app_company" VARCHAR NOT NULL,
  "app_version" VARCHAR NOT NULL,
  "app_status" INTEGER DEFAULT 1,
  "app_data" TEXT DEFAULT NULL,
  "created_at" INTEGER DEFAULT 0,
  "updated_at" INTEGER DEFAULT 0,
  "deleted_at" INTEGER DEFAULT 0
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS rules;
-- +goose StatementEnd
