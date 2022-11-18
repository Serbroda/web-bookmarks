-- +goose Up
-- +goose StatementBegin
CREATE TABLE roles (
    id int NOT NULL AUTO_INCREMENT,
    created_at timestamp default CURRENT_TIMESTAMP,
    updated_at timestamp,
    deleted_at timestamp,
    name varchar(80) NOT NULL,
    CONSTRAINT PK_roles PRIMARY KEY (id),
    CONSTRAINT UC_roles_name UNIQUE (name)
);
-- +goose StatementEnd
-- +goose StatementBegin
INSERT INTO roles(created_at, name)
VALUES(CURRENT_TIMESTAMP, 'Admin');
-- +goose StatementEnd
-- +goose StatementBegin
INSERT INTO roles(created_at, name)
VALUES(CURRENT_TIMESTAMP, 'User');
-- +goose StatementEnd
-- +goose StatementBegin
INSERT INTO roles(created_at, name)
VALUES(CURRENT_TIMESTAMP, 'Owner');
-- +goose StatementEnd
-- +goose StatementBegin
INSERT INTO roles(created_at, name)
VALUES(CURRENT_TIMESTAMP, 'Maintainer');
-- +goose StatementEnd
-- +goose StatementBegin
INSERT INTO roles(created_at, name)
VALUES(CURRENT_TIMESTAMP, 'Guest');
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE roles;
-- +goose StatementEnd