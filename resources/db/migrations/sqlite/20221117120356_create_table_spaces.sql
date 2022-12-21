-- +goose Up
-- +goose StatementBegin
CREATE TABLE spaces (
                        id INTEGER PRIMARY KEY AUTOINCREMENT,
                        created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
                        updated_at DATETIME,
                        deleted_at DATETIME,
                        short_id TEXT NOT NULL,
                        owner_id INTEGER NOT NULL,
                        name TEXT NOT NULL,
                        description TEXT,
                        visibility TEXT NOT NULL,
                        UNIQUE (short_id),
                        FOREIGN KEY (owner_id) REFERENCES users(id)
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE users_spaces (
                              user_id INTEGER NOT NULL,
                              space_id INTEGER NOT NULL,
                              role_id INTEGER NOT NULL,
                              created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
                              PRIMARY KEY (user_id, space_id),
                              UNIQUE (user_id, space_id),
                              FOREIGN KEY (user_id) REFERENCES users(id),
                              FOREIGN KEY (space_id) REFERENCES spaces(id),
                              FOREIGN KEY (role_id) REFERENCES roles(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE spaces;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE users_spaces;
-- +goose StatementEnd