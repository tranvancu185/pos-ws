-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS payments (
  "payment_id" INTEGER PRIMARY KEY AUTOINCREMENT,
  "payment_order_id" INTEGER NOT NULL,
  "payment_order_code" VARCHAR(16) NOT NULL,
  "payment_amount" INTEGER NOT NULL DEFAULT 0,
  "payment_method" INTEGER NOT NULL DEFAULT 1,
  "payment_status" INTEGER NOT NULL DEFAULT 1,
  "payment_properties" TEXT DEFAULT NULL,
  "created_at" INTEGER DEFAULT 0,
  "updated_at" INTEGER DEFAULT 0,
  "deleted_at" INTEGER DEFAULT 0
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
