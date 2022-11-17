-- +goose Up
-- +goose StatementBegin
CREATE TABLE spaces (
    id bigint NOT NULL AUTO_INCREMENT,
    created_at timestamp default CURRENT_TIMESTAMP,
    updated_at timestamp,
    deleted_at timestamp,
    short_id varchar(20),
    owner_id bigint NOT NULL,
    name varchar(40),
    description varchar(255),
    visibility varchar(20),
    CONSTRAINT PK_spaces PRIMARY KEY (id),
    CONSTRAINT UC_spaces_short_id UNIQUE (short_id),
    CONSTRAINT FK_spaces_owner_id FOREIGN KEY (owner_id) REFERENCES users(id)
);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE spaces;
-- +goose StatementEnd