-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS rules (
  "rule_id" INTEGER PRIMARY KEY AUTOINCREMENT,
  "rule_name" VARCHAR(30) NOT NULL,
  "rule_code" VARCHAR(5) NOT NULL,
  "rule_type" INTEGER NOT NULL DEFAULT 0,
  "rule_status" INTEGER NOT NULL DEFAULT 1,
  "rule_properties" TEXT DEFAULT NULL,
  "created_at" INTEGER DEFAULT 0,
  "updated_at" INTEGER DEFAULT 0,
  "deleted_at" INTEGER DEFAULT 0
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS rules;
-- +goose StatementEnd
