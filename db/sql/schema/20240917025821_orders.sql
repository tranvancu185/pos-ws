-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS orders (
  "order_id" INTEGER PRIMARY KEY AUTOINCREMENT,
  "order_code" VARCHAR(16) NOT NULL UNIQUE,
  "order_create_at" INTEGER NOT NULL DEFAULT 0,
  "order_update_at" INTEGER NOT NULL DEFAULT 0,
  "order_status" INTEGER NOT NULL DEFAULT 1,
  "order_total_amount" INTEGER NOT NULL DEFAULT 0,
  "order_total_discount" INTEGER NOT NULL DEFAULT 0,
  "order_properties" TEXT DEFAULT NULL,
  "order_balance" INTEGER NOT NULL DEFAULT 0,
  "order_user_id" INTEGER NOT NULL DEFAULT 0,
  "order_table_id" INTEGER NOT NULL DEFAULT 0,
  "order_customer_id" INTEGER NOT NULL DEFAULT 0
);

CREATE INDEX idx_order_code ON orders (order_code);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS orders;
-- +goose StatementEnd
