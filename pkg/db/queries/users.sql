-- name: FindUser :one
SELECT *
FROM users u
WHERE id = ?
LIMIT 1;
-- name: FindUserByUsername :one
SELECT *
FROM users
WHERE lower(username) = lower(?)
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
-- name: FindRolesByUser :many
SELECT r.*
from roles r
    inner join users_roles ur on ur.role_id = r.id
where ur.user_id = ?;
-- name: InsertUserRole :exec
INSERT INTO users_roles (user_id, role_id, created_at)
VALUES (?, ?, CURRENT_TIMESTAMP);