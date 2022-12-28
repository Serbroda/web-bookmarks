-- +goose Up
-- +goose StatementBegin
CREATE TABLE roles
(
    id          INTEGER PRIMARY KEY AUTOINCREMENT,
    created_at  DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at  DATETIME,
    deleted_at  DATETIME,
    name        TEXT NOT NULL,
    description TEXT,
    UNIQUE (name)
);
-- +goose StatementEnd

-- +goose StatementBegin
INSERT INTO roles(created_at, name, description)
VALUES(CURRENT_TIMESTAMP, 'ADMIN', 'Administrator');
-- +goose StatementEnd
-- +goose StatementBegin
INSERT INTO roles(created_at, name, description)
VALUES(CURRENT_TIMESTAMP, 'USER', 'User');
-- +goose StatementEnd
-- +goose StatementBegin
INSERT INTO roles(created_at, name, description)
VALUES(CURRENT_TIMESTAMP, 'OWNER', 'Owner');
-- +goose StatementEnd
-- +goose StatementBegin
INSERT INTO roles(created_at, name, description)
VALUES(CURRENT_TIMESTAMP, 'MAINTAINER', 'Maintainer');
-- +goose StatementEnd
-- +goose StatementBegin
INSERT INTO roles(created_at, name, description)
VALUES(CURRENT_TIMESTAMP, 'GUEST', 'Guest');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE roles;
-- +goose StatementEnd