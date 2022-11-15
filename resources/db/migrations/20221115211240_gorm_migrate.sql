-- +goose Up
-- +goose StatementBegin
-- users definition
CREATE TABLE IF NOT EXISTS users (
    id integer,
    created_at datetime,
    updated_at datetime,
    deleted_at datetime,
    username text UNIQUE,
    password text,
    name text,
    email text,
    PRIMARY KEY (id)
);
CREATE TABLE IF NOT EXISTS groups (
    id text,
    created_at datetime,
    updated_at datetime,
    deleted_at datetime,
    owner_id integer,
    icon text,
    name text,
    description text,
    visibility text,
    PRIMARY KEY (id),
    CONSTRAINT fk_groups_owner FOREIGN KEY (owner_id) REFERENCES users(id)
);
CREATE TABLE IF NOT EXISTS links (
    id text,
    created_at datetime,
    updated_at datetime,
    deleted_at datetime,
    group_id text,
    name text,
    url text,
    description text,
    visibility text,
    PRIMARY KEY (id),
    CONSTRAINT fk_groups_links FOREIGN KEY (group_id) REFERENCES groups(id)
);
CREATE TABLE IF NOT EXISTS group_subscriptions (
    id integer,
    created_at datetime,
    updated_at datetime,
    deleted_at datetime,
    user_id integer,
    group_id text,
    PRIMARY KEY (id),
    CONSTRAINT fk_group_subscriptions_user FOREIGN KEY (user_id) REFERENCES users(id),
    CONSTRAINT fk_group_subscriptions_group FOREIGN KEY (group_id) REFERENCES groups(id)
);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
DROP TABLE groups;
DROP TABLE links;
DROP TABLE group_subscriptions;
-- +goose StatementEnd