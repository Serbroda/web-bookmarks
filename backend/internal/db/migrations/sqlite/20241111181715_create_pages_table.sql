-- +goose Up
-- +goose StatementBegin
CREATE TABLE pages
(
    id          INTEGER PRIMARY KEY AUTOINCREMENT,
    space_id    INTEGER NOT NULL,
    name        TEXT    NOT NULL,
    description TEXT,
    visibility  TEXT    NOT NULL DEFAULT 'PRIVATE',
    parent_id   INTEGER,
    created_at  DATETIME         DEFAULT CURRENT_TIMESTAMP,
    updated_at  DATETIME         DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_pages_space_id FOREIGN KEY (space_id) REFERENCES spaces (id),
    CONSTRAINT fk_pages_parent_id FOREIGN KEY (parent_id) REFERENCES pages (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE pages;
-- +goose StatementEnd
