-- +goose Up
-- +goose StatementBegin
CREATE TABLE links (
    id bigint NOT NULL AUTO_INCREMENT,
    created_at timestamp default CURRENT_TIMESTAMP,
    updated_at timestamp,
    deleted_at timestamp,
    short_id varchar(20),
    page_id bigint NOT NULL,
    name varchar(128),
    url varchar(2048) NOT NULL,
    description varchar(2048),
    visibility varchar(20),
    CONSTRAINT PK_links PRIMARY KEY (id),
    CONSTRAINT UC_links_short_id UNIQUE (short_id),
    CONSTRAINT FK_links_page_id FOREIGN KEY (page_id) REFERENCES pages(id)
);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE links;
-- +goose StatementEnd