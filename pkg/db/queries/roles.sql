-- name: FindRoleByName :one
SELECT *
FROM roles
where lower(name) = lower(?)
LIMIT 1;
-- name: FindRolesByNames :many
SELECT *
FROM roles
where name in (sqlc.slice('names'));
-- name: CreateRole :execlastid
INSERT INTO roles(created_at, name, description)
VALUES(CURRENT_TIMESTAMP, ?, ?);