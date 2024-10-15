-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS products (
  "product_id" INTEGER PRIMARY KEY AUTOINCREMENT,
  "product_code" VARCHAR(6) NOT NULL UNIQUE,
  "product_name" VARCHAR(255) NOT NULL,
  "product_display_name" VARCHAR(255) NOT NULL,
  "product_description" TEXT DEFAULT NULL,
  "product_price" INTEGER NOT NULL DEFAULT 0,
  "product_status" INTEGER NOT NULL DEFAULT 1,
  "product_properties" TEXT DEFAULT NULL,
  "product_category_id" INTEGER NOT NULL DEFAULT 0,
  "created_at" INTEGER DEFAULT 0,
  "updated_at" INTEGER DEFAULT 0,
  "deleted_at" INTEGER DEFAULT 0
);
CREATE UNIQUE INDEX idx_product_code ON products (product_code);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
