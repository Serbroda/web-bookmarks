-- +goose Up
-- +goose StatementBegin
CREATE TABLE pages (
    id bigint NOT NULL AUTO_INCREMENT,
    created_at timestamp default CURRENT_TIMESTAMP,
    updated_at timestamp,
    deleted_at timestamp,
    short_id varchar(20) NOT NULL,
    owner_id bigint NOT NULL,
    space_id bigint NOT NULL,
    parent_id bigint,
    name varchar(128) NOT NULL,
    description varchar(2048),
    visibility varchar(20) NOT NULL,
    CONSTRAINT PK_pages PRIMARY KEY (id),
    CONSTRAINT UC_pages_short_id UNIQUE INDEX (short_id),
    CONSTRAINT FK_pages_owner_id FOREIGN KEY (owner_id) REFERENCES users(id),
    CONSTRAINT FK_pages_space_id FOREIGN KEY (space_id) REFERENCES spaces(id),
    CONSTRAINT FK_pages_parent_id FOREIGN KEY (parent_id) REFERENCES pages(id)
);
-- +goose StatementEnd
-- +goose StatementBegin
CREATE TABLE users_pages (
    user_id bigint NOT NULL,
    page_id bigint NOT NULL,
    role_id int NOT NULL,
    created_at timestamp default CURRENT_TIMESTAMP,
    CONSTRAINT PK_users_pages PRIMARY KEY (user_id, page_id),
    CONSTRAINT UC_users_pages_user_id_page_id UNIQUE (user_id, page_id),
    CONSTRAINT FK_users_pages_user_id FOREIGN KEY (user_id) REFERENCES users(id),
    CONSTRAINT FK_users_pages_page_id FOREIGN KEY (page_id) REFERENCES pages(id),
    CONSTRAINT FK_users_pages_role_id FOREIGN KEY (role_id) REFERENCES roles(id)
);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE pages;
-- +goose StatementEnd
-- +goose StatementBegin
DROP TABLE users_pages;
-- +goose StatementEnd