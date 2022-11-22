-- name: FindAllRoles :many
SELECT *
FROM roles;
-- name: FindRoleByName :one
SELECT *
FROM roles
WHERE lower(name) = lower(?)
LIMIT 1;
-- name: InsertUserRole :exec
INSERT INTO users_roles(user_id, role_id)
VALUES(?, ?);
-- name: DeleteUserRole :exec
DELETE FROM users_roles
WHERE user_id = ?
    AND role_id = ?;
-- name: CountUserRole :one
SELECT count(*)
FROM users_roles ur
    LEFT JOIN roles r on r.id = ur.role_id
WHERE ur.user_id = ?
    AND r.name = ?;