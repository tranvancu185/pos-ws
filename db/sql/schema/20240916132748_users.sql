-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
  "user_id" INTEGER PRIMARY KEY AUTOINCREMENT,
  "user_display_name" VARCHAR(50) NOT NULL,
  "user_name" VARCHAR(30) NOT NULL UNIQUE,
  "user_phone" VARCHAR(15) NOT NULL UNIQUE,
  "user_password" VARCHAR(32) NOT NULL,
  "user_status" INTEGER NOT NULL DEFAULT 1,
  "user_avatar" VARCHAR NOT NULL DEFAULT '',
  "user_role_id" INTEGER NOT NULL DEFAULT 0,
  "created_at" INTEGER NOT NULL DEFAULT 0,
  "updated_at" INTEGER NOT NULL DEFAULT 0,
  "deleted_at" INTEGER NOT NULL DEFAULT 0,
  "last_login_at" INTEGER NOT NULL DEFAULT 0,
  "login_times" INTEGER NOT NULL DEFAULT 0
);

CREATE UNIQUE INDEX idx_user_name ON users (user_name);
CREATE UNIQUE INDEX idx_user_phone ON users (user_phone);

INSERT INTO users (user_display_name, user_name, user_phone, user_password, user_status, user_role_id, created_at, updated_at, deleted_at, last_login_at, login_times) VALUES ('admin Vey', 'admin', '0965135278', '8c6976e5b5410415bde908bd4dee15dfb167a9c873fc4bb8a81f6f2ab448a918', 1, 1, 1631788068, 1631788068, 0, 0, 0);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `users`;
-- +goose StatementEnd
