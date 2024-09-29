-- name: FindUserById :one
SELECT *
FROM users u
WHERE id = ? LIMIT 1;

-- name: FindUserByEmail :one
SELECT *
FROM users u
WHERE lower(email) = ? LIMIT 1;

-- name: FindUserByUsername :one
SELECT *
FROM users u
WHERE lower(username) = ? LIMIT 1;

-- name: CreateUser :execlastid
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