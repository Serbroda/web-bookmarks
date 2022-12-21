-- +goose Up
-- +goose StatementBegin
CREATE TABLE spaces (
    id bigint NOT NULL AUTO_INCREMENT,
    created_at timestamp default CURRENT_TIMESTAMP,
    updated_at timestamp,
    deleted_at timestamp,
    short_id varchar(20) NOT NULL,
    owner_id bigint NOT NULL,
    name varchar(128) NOT NULL,
    description varchar(2048),
    visibility varchar(20) NOT NULL,
    CONSTRAINT PK_spaces PRIMARY KEY (id),
    CONSTRAINT UC_spaces_short_id UNIQUE INDEX (short_id),
    CONSTRAINT FK_spaces_owner_id FOREIGN KEY (owner_id) REFERENCES users(id)
);
-- +goose StatementEnd
-- +goose StatementBegin
CREATE TABLE users_spaces (
    user_id bigint NOT NULL,
    space_id bigint NOT NULL,
    role_id int NOT NULL,
    created_at timestamp default CURRENT_TIMESTAMP,
    CONSTRAINT PK_users_spaces PRIMARY KEY (user_id, space_id),
    CONSTRAINT UC_users_spaces_user_id_space_id UNIQUE (user_id, space_id),
    CONSTRAINT FK_users_spaces_user_id FOREIGN KEY (user_id) REFERENCES users(id),
    CONSTRAINT FK_users_spaces_space_id FOREIGN KEY (space_id) REFERENCES spaces(id),
    CONSTRAINT FK_users_spaces_role_id FOREIGN KEY (role_id) REFERENCES roles(id)
);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE spaces;
-- +goose StatementEnd
-- +goose StatementBegin
DROP TABLE users_spaces;
-- +goose StatementEnd