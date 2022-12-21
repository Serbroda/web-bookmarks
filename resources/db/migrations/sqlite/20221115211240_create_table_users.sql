-- +goose Up
-- +goose StatementBegin
CREATE TABLE users
(
    id                      INTEGER PRIMARY KEY AUTOINCREMENT,
    created_at              DATETIME         DEFAULT CURRENT_TIMESTAMP,
    updated_at              DATETIME,
    deleted_at              DATETIME,
    first_name              TEXT    NOT NULL,
    last_name               TEXT    NOT NULL,
    username                TEXT    NOT NULL,
    password                TEXT    NOT NULL,
    email                   TEXT    NOT NULL,
    active                  INTEGER NOT NULL DEFAULT 0,
    activation_confirmed_at DATETIME,
    UNIQUE (username)
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE users_roles
(
    user_id    INTEGER NOT NULL,
    role_id    INTEGER NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id, role_id),
    UNIQUE (user_id, role_id),
    FOREIGN KEY (user_id) REFERENCES users (id),
    FOREIGN KEY (role_id) REFERENCES roles (id)
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE activation_tokens
(
    user_id    INTEGER NOT NULL,
    token_hash TEXT    NOT NULL,
    expires_at DATETIME,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id, token_hash)
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE password_reset_tokens
(
    user_id    INTEGER  NOT NULL,
    token_hash TEXT     NOT NULL,
    expires_at DATETIME NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id, token_hash)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE users_roles;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE activation_tokens;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE password_reset_tokens;
-- +goose StatementEnd