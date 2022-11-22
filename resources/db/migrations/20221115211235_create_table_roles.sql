-- +goose Up
-- +goose StatementBegin
CREATE TABLE roles (
    id int NOT NULL AUTO_INCREMENT,
    created_at timestamp default CURRENT_TIMESTAMP,
    updated_at timestamp,
    deleted_at timestamp,
    name varchar(80) NOT NULL,
    description varchar(255),
    CONSTRAINT PK_roles PRIMARY KEY (id),
    CONSTRAINT UC_roles_name UNIQUE (name)
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