-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS customers (
  "customer_id" INTEGER PRIMARY KEY AUTOINCREMENT,
  "customer_name" VARCHAR(255) NOT NULL,
  "customer_email" VARCHAR(255) DEFAULT NULL,
  "customer_phone" VARCHAR(30) NOT NULL UNIQUE,
  "customer_status" INTEGER NOT NULL DEFAULT 1,
  "customer_total_orders" INTEGER NOT NULL DEFAULT 0,
  "customer_properties" TEXT DEFAULT NULL,
  "created_at" INTEGER DEFAULT 0,
  "updated_at" INTEGER DEFAULT 0,
  "deleted_at" INTEGER DEFAULT 0
);
CREATE UNIQUE INDEX idx__customer_phone ON customers (customer_phone);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
