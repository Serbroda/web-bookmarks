-- +goose Up
-- +goose StatementBegin
-- users definition
CREATE TABLE users (
    id bigint NOT NULL AUTO_INCREMENT,
    created_at timestamp default CURRENT_TIMESTAMP,
    updated_at timestamp,
    deleted_at timestamp,
    username varchar(120) NOT NULL,
    password varchar(120),
    name varchar(120),
    email varchar(120),
    CONSTRAINT PK_users PRIMARY KEY (id),
    CONSTRAINT UC_users_username UNIQUE (username)
);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd