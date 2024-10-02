-- +goose Up
-- +goose StatementBegin
CREATE TABLE spaces
(
    id          INTEGER PRIMARY KEY AUTOINCREMENT,
    name        TEXT NOT NULL,
    description TEXT,
    visibility  TEXT NOT NULL DEFAULT 'PRIVATE',
    created_at  DATETIME      DEFAULT CURRENT_TIMESTAMP,
    updated_at  DATETIME      DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE spaces_users
(
    space_id   INTEGER NOT NULL,
    user_id    INTEGER NOT NULL,
    role       TEXT    NOT NULL DEFAULT 'MEMBER',
    created_at DATETIME         DEFAULT CURRENT_TIMESTAMP,
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
