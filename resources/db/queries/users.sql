-- name: FindUser :one
SELECT *
FROM users
WHERE id = ?
LIMIT 1;
-- name: FindUserByName :one
SELECT *
FROM users
WHERE lower(username) = lower(?)
LIMIT 1;
-- name: CountUserByName :one
SELECT count(*)
FROM users
WHERE lower(username) = lower(?);
-- name: CreateUser :execlastid
INSERT INTO users (
        created_at,
        username,
        password,
        name,
        email,
        active,
        must_change_password
    )
VALUES(
        CURRENT_TIMESTAMP,
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
-- name: DeleteUserFull :exec
DELETE FROM users
WHERE id = ?;