-- name: FindUserById :one
SELECT *
FROM users u
WHERE id = ? LIMIT 1;

-- name: FindUserByEmailOrUsername :one
SELECT *
FROM users u
WHERE lower(email) = lower(sqlc.arg('email'))
   or lower(username) = lower(sqlc.arg('username')) LIMIT 1;

-- name: CountUserByEmail :one
SELECT count(*)
FROM users
WHERE lower(email) = lower(sqlc.arg('email')) LIMIT 1;

-- name: CountUserByUsername :one
SELECT count(*)
FROM users
WHERE lower(username) = lower(sqlc.arg('username')) LIMIT 1;

-- name: CreateUser :execlastid
INSERT INTO users (created_at,
                   updated_at,
                   email,
                   username,
                   password,
                   display_name)
VALUES (CURRENT_TIMESTAMP,
        CURRENT_TIMESTAMP,
        lower(sqlc.arg('email')),
        lower(sqlc.arg('username')),
        sqlc.arg('password'),
        sqlc.arg('display_name'));

-- name: UpdatePassword :exec
UPDATE users
SET password = ?
WHERE id = ?
;

