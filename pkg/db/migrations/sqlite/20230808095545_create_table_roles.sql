-- +goose Up
-- +goose StatementBegin
CREATE TABLE roles (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME,
    deleted_at DATETIME,
    name TEXT NOT NULL UNIQUE,
    description TEXT
);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE roles;
-- +goose StatementEnd