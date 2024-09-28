-- name: FindById :one
SELECT *
FROM users u
WHERE id = ? LIMIT 1;

-- name: FindByEmail :one
SELECT *
FROM users u
WHERE lower(username) = ? LIMIT 1;

-- name: FindByUsername :one
SELECT *
FROM users u
WHERE lower(username) = ? LIMIT 1;

-- name: Create :execlastid
INSERT INTO users (created_at,
                   updated_at,
                   email,
                   username,
                   password)
VALUES (CURRENT_TIMESTAMP,
        CURRENT_TIMESTAMP,
        ?,
        ?,
        ?);