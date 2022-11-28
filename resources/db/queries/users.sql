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
-- name: FindByEmail :one
SELECT *
FROM users
WHERE lower(email) = lower(?)
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
    activation_confirmed_at = COALESCE(?, activation_confirmed_at)
WHERE id = ?;
-- name: DeleteUser :exec
DELETE FROM users
WHERE id = ?;
-- name: FindActivationCode :one
SELECT *
FROM activation_tokens
WHERE token_hash = ?
LIMIT 1;
-- name: InsertActivationToken :exec
INSERT INTO activation_tokens (user_id, token_hash, expires_at, created_at)
VALUES(?, ?, ?, CURRENT_TIMESTAMP);
-- name: FindPasswordResetCodeByUserIdAndToken :one
SELECT *
FROM password_reset_tokens
WHERE user_id = ?
    and token_hash = ?
LIMIT 1;
-- name: FindPasswordResetCodeByEmailAndToken :one
SELECT prt.*,
    u.active as UserActive
FROM password_reset_tokens prt
    INNER JOIN users u on u.id = prt.user_id
WHERE u.email = ?
    and token_hash = ?
LIMIT 1;
-- name: InsertPasswordResetToken :exec
INSERT INTO password_reset_tokens (user_id, token_hash, expires_at, created_at)
VALUES(?, ?, ?, CURRENT_TIMESTAMP);