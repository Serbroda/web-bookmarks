-- name: FindUser :one
SELECT *
FROM users
WHERE id = ?
LIMIT 1;
-- name: FindUserByUsername :one
SELECT *
FROM users
WHERE lower(username) = lower(?)
LIMIT 1;
-- name: FindUserByActivationCode :one
SELECT *
FROM users
WHERE activation_code = ?
LIMIT 1;
-- name: CreateUser :execlastid
INSERT INTO users (
        created_at,
        username,
        password,
        email,
        first_name,
        last_name,
        active,
        activation_code,
        activation_sent_at,
        activation_code_expires_at,
        activation_confirmed_at
    )
VALUES(
        CURRENT_TIMESTAMP,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?
    );
-- name: UpdateUser :exec
UPDATE users
SET updated_at = CURRENT_TIMESTAMP,
    password = COALESCE(?, password),
    email = COALESCE(?, email),
    first_name = COALESCE(?, first_name),
    last_name = COALESCE(?, last_name),
    active = COALESCE(?, active),
    activation_code = COALESCE(?, activation_code),
    activation_sent_at = COALESCE(?, activation_sent_at),
    activation_code_expires_at = COALESCE(?, activation_sent_at),
    activation_confirmed_at = COALESCE(?, activation_confirmed_at)
WHERE id = ?;
-- name: DeleteUser :exec
DELETE FROM users
WHERE id = ?;