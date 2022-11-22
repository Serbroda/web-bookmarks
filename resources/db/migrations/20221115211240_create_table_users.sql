-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    id bigint NOT NULL AUTO_INCREMENT,
    created_at timestamp default CURRENT_TIMESTAMP,
    updated_at timestamp,
    deleted_at timestamp,
    username varchar(120) NOT NULL,
    password varchar(120) NOT NULL,
    name varchar(120),
    email varchar(120) NOT NULL,
    active BOOLEAN NOT NULL default FALSE,
    confirmed_at timestamp,
    must_change_password BOOLEAN NOT NULL default FALSE,
    CONSTRAINT PK_users PRIMARY KEY (id),
    CONSTRAINT UC_users_username UNIQUE (username)
);
-- +goose StatementEnd
-- +goose StatementBegin
CREATE TABLE users_roles (
    user_id bigint NOT NULL,
    role_id int NOT NULL,
    created_at timestamp default CURRENT_TIMESTAMP,
    CONSTRAINT PK_users_roles PRIMARY KEY (user_id, role_id),
    CONSTRAINT UC_users_roles_user_id_role_id UNIQUE (user_id, role_id),
    CONSTRAINT FK_users_roles_user_id FOREIGN KEY (user_id) REFERENCES users(id),
    CONSTRAINT FK_users_roles_role_id FOREIGN KEY (role_id) REFERENCES roles(id)
);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
DROP TABLE users_roles;
-- +goose StatementEnd