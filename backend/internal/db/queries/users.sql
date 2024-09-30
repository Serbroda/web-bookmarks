-- name: FindUserById :one
SELECT *
FROM users u
WHERE id = ? LIMIT 1;

-- name: FindUserByEmail :one
SELECT *
FROM users u
WHERE lower(email) = lower(?) LIMIT 1;

-- name: CountByUsernameAndTag :one
SELECT count(*)
FROM users
WHERE lower(username) = lower(?)
  AND tag = ? LIMIT 1;

-- name: FindUserByUsername :one
SELECT *
FROM users u
WHERE lower(username) = ?
  AND tag = ? LIMIT 1;


-- name: CreateUser :execlastid
INSERT INTO users (created_at,
                   updated_at,
                   email,
                   username,
                   tag,
                   password)
VALUES (CURRENT_TIMESTAMP,
        CURRENT_TIMESTAMP,
        ?,
        ?,
        ?,
        ?);

-- name: UpdateUsername :exec
UPDATE users
SET username = ?,
    tag      = ?
;

-- name: UpdatePassword :exec
UPDATE users
SET username = ?,
    tag      = ?
;

-- name: UpdateEmail :exec
UPDATE users
SET username = ?,
    tag      = ?
;
