-- name: FindPageById :one
SELECT *
FROM pages p
WHERE id = ? LIMIT 1;

-- name: FindRootPagesBySpaceId :many
SELECT *
FROM pages p
WHERE space_id = ? AND parent_id IS NULL;

-- name: FindPagesBySpaceId :many
SELECT *
FROM pages p
WHERE space_id = ?;