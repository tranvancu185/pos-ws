-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS counters (
    "counter_id" INTEGER PRIMARY KEY AUTOINCREMENT,
    "counter_name" VARCHAR(50) NOT NULL,
    "counter_number" INTEGER NOT NULL DEFAULT 0,
    "created_at" INTEGER DEFAULT 0,
    "updated_at" INTEGER DEFAULT 0
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS counters;
-- +goose StatementEnd
