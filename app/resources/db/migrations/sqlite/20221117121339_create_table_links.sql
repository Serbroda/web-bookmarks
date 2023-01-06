-- +goose Up
-- +goose StatementBegin
CREATE TABLE links
(
    id          INTEGER PRIMARY KEY AUTOINCREMENT,
    created_at  DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at  DATETIME,
    deleted_at  DATETIME,
    short_id    TEXT,
    page_id     INTEGER NOT NULL,
    name        TEXT,
    url         TEXT    NOT NULL,
    description TEXT,
    visibility  TEXT,
    UNIQUE (short_id),
    FOREIGN KEY (page_id) REFERENCES pages (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE links;
-- +goose StatementEnd