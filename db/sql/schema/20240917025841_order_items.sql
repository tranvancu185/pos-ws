-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS order_items (
  "odi_id" INTEGER PRIMARY KEY AUTOINCREMENT,
  "odi_order_id" INTEGER NOT NULL,
  "odi_product_code" INTEGER NOT NULL,
  "odi_quantity" INTEGER NOT NULL DEFAULT 0,
  "odi_price" INTEGER NOT NULL DEFAULT 0,
  "odi_discount" INTEGER NOT NULL DEFAULT 0,
  FOREIGN KEY ("odi_order_id") REFERENCES orders ("odi_order_id"),
  FOREIGN KEY ("odi_product_code") REFERENCES products ("odi_product_code")
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS order_items;
-- +goose StatementEnd
