-- +goose Up
-- +goose StatementBegin
CREATE TABLE pages (
    id bigint NOT NULL AUTO_INCREMENT,
    created_at timestamp default CURRENT_TIMESTAMP,
    updated_at timestamp,
    deleted_at timestamp,
    short_id varchar(20),
    owner_id bigint NOT NULL,
    space_id bigint NOT NULL,
    parent_id bigint,
    name varchar(40),
    description varchar(255),
    visibility varchar(20),
    CONSTRAINT PK_pages PRIMARY KEY (id),
    CONSTRAINT UC_pages_short_id UNIQUE (short_id),
    CONSTRAINT FK_pages_owner_id FOREIGN KEY (owner_id) REFERENCES users(id),
    CONSTRAINT FK_pages_space_id FOREIGN KEY (space_id) REFERENCES spaces(id),
    CONSTRAINT FK_pages_parent_id FOREIGN KEY (parent_id) REFERENCES pages(id)
);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE pages;
-- +goose StatementEnd