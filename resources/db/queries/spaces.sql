-- name: GetSpace :one
SELECT *
FROM spaces
WHERE short_id = ?
    AND deleted_at IS NULL
LIMIT 1;
-- name: ListSpaces :many
SELECT *
FROM spaces
WHERE owner_id = ?
    AND deleted_at IS NULL;
-- name: DeleteSpace :exec
UPDATE spaces
SET deleted_at = CURRENT_TIMESTAMP
WHERE short_id = ?
    AND deleted_at IS NULL;