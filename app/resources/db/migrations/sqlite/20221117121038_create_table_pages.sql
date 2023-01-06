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

-- +goose StatementBegin
CREATE TABLE users_pages
(
    user_id    INTEGER NOT NULL,
    page_id    INTEGER NOT NULL,
    role_id    INTEGER NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id, page_id),
    UNIQUE (user_id, page_id),
    FOREIGN KEY (user_id) REFERENCES users (id),
    FOREIGN KEY (page_id) REFERENCES pages (id),
    FOREIGN KEY (role_id) REFERENCES roles (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE pages;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE users_pages;
-- +goose StatementEnd