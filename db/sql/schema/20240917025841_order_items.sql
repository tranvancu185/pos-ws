-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS order_detail (
  "odt_id" INTEGER PRIMARY KEY AUTOINCREMENT,
  "odt_order_id" INTEGER NOT NULL,
  "odt_product_code" INTEGER NOT NULL,
  "odt_quantity" INTEGER NOT NULL DEFAULT 0,
  "odt_price" INTEGER NOT NULL DEFAULT 0,
  "odt_discount" INTEGER NOT NULL DEFAULT 0,
  "odt_properties" TEXT DEFAULT NULL,
  FOREIGN KEY ("odt_order_id") REFERENCES orders ("order_id"),
  FOREIGN KEY ("odt_product_code") REFERENCES products ("product_code")
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS order_detail;
-- +goose StatementEnd
