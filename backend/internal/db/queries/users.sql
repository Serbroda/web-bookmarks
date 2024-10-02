-- name: FindUserById :one
SELECT *
FROM users u
WHERE id = ? LIMIT 1;

-- name: FindUserByEmailOrUsername :one
SELECT *
FROM users u
WHERE lower(email) = ?
   or lower(username) = ? LIMIT 1;

-- name: CountUserByEmail :one
SELECT count(*)
FROM users
WHERE lower(email) = ? LIMIT 1;

-- name: CountUserByUsername :one
SELECT count(*)
FROM users
WHERE lower(username) = ? LIMIT 1;

-- name: CreateUser :execlastid
INSERT INTO users (created_at,
                   updated_at,
                   email,
                   username,
                   password,
                   display_name)
VALUES (CURRENT_TIMESTAMP,
        CURRENT_TIMESTAMP,
        ?,
        ?,
        ?,
        ?);

-- name: UpdatePassword :exec
UPDATE users
SET password = ?
WHERE id = ?
;

