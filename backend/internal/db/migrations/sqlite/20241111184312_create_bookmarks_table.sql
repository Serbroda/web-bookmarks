-- +goose Up
-- +goose StatementBegin
CREATE TABLE bookmarks
(
    id          INTEGER PRIMARY KEY AUTOINCREMENT,
    page_id     INTEGER NOT NULL,
    name        TEXT    NOT NULL,
    description TEXT,
    url         TEST    NOT NULL,
    created_at  DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at  DATETIME DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_pages_page_id FOREIGN KEY (page_id) REFERENCES spaces (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE bookmarks;
-- +goose StatementEnd
