-- +goose Up
-- +goose StatementBegin
CREATE TABLE spaces
(
    id          INTEGER PRIMARY KEY AUTOINCREMENT,
    name        TEXT    NOT NULL,
    description TEXT,
    owner_id    INTEGER NOT NULL,
    visibility  TEXT    NOT NULL DEFAULT 'PRIVATE',
    created_at  DATETIME         DEFAULT CURRENT_TIMESTAMP,
    updated_at  DATETIME         DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_spaces_owner_id FOREIGN KEY (owner_id) REFERENCES users (id)
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE users_spaces
(
    user_id    INTEGER NOT NULL,
    space_id   INTEGER NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_users_spaces_user_id FOREIGN KEY (user_id) REFERENCES users (id),
    CONSTRAINT fk_users_spaces_space_id FOREIGN KEY (space_id) REFERENCES spaces (id),
    PRIMARY KEY (user_id, space_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE spaces;
-- +goose StatementEnd
-- +goose StatementBegin
DROP TABLE users_spaces;
-- +goose StatementEnd
