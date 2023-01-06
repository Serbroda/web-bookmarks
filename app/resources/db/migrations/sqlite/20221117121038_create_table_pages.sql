-- +goose Up
-- +goose StatementBegin
CREATE TABLE pages
(
    id          INTEGER PRIMARY KEY AUTOINCREMENT,
    created_at  DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at  DATETIME,
    deleted_at  DATETIME,
    short_id    TEXT    NOT NULL,
    owner_id    INTEGER NOT NULL,
    space_id    INTEGER NOT NULL,
    parent_id   INTEGER,
    name        TEXT    NOT NULL,
    description TEXT,
    visibility  TEXT    NOT NULL,
    UNIQUE (short_id),
    FOREIGN KEY (owner_id) REFERENCES users (id),
    FOREIGN KEY (space_id) REFERENCES spaces (id),
    FOREIGN KEY (parent_id) REFERENCES pages (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE pages;
-- +goose StatementEnd
