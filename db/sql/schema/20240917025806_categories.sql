-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS categories (
  "category_id" INTEGER PRIMARY KEY AUTOINCREMENT,
  "category_name" VARCHAR(255) NOT NULL,
  "category_status" INTEGER NOT NULL DEFAULT 1,
  "category_description" TEXT DEFAULT NULL,
  "category_properties" TEXT DEFAULT NULL,
  "created_at" INTEGER DEFAULT 0,
  "updated_at" INTEGER DEFAULT 0,
  "deleted_at" INTEGER DEFAULT 0
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS categories;
-- +goose StatementEnd
