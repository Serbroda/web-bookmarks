-- name: GetUser :one
SELECT *
FROM users
WHERE id = ?
LIMIT 1;
-- name: GetUserByUsername :one
SELECT *
FROM users
WHERE lower(username) = lower(?)
LIMIT 1;
-- name: ExistsUser :one
SELECT count(*)
FROM users
WHERE lower(username) = lower(?);
-- name: CreateUser :execlastid
INSERT INTO users (
        created_at,
        username,
        name,
        password,
        email
    )
VALUES(
        CURRENT_TIMESTAMP,
        ?,
        ?,
        ?,
        ?
    );
-- name: UpdateUser :exec
UPDATE users
SET updated_at = CURRENT_TIMESTAMP,
    password = ?,
    name = ?,
    email = ?
WHERE id = ?
    AND deleted_at IS NULL;
-- name: DeleteUser :exec
UPDATE users
SET deleted_at = CURRENT_TIMESTAMP
WHERE id = ?
    AND deleted_at IS NULL;