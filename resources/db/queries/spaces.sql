-- name: GetSpace :one
SELECT *
FROM spaces
WHERE short_id = ?
LIMIT 1;
-- name: ListSpaces :many
SELECT *
FROM spaces
WHERE owner_id = ?
    AND deleted_at IS NULL;
-- name: UpdateSpace :exec
UPDATE pages
SET updated_at = CURRENT_TIMESTAMP,
    owner_id = ?,
    space_id = ?,
    parent_id = ?,
    name = ?,
    description = ?,
    visibility = ?
WHERE short_id = 0;
-- name: DeleteSpace :exec
UPDATE spaces
SET deleted_at = CURRENT_TIMESTAMP
WHERE short_id = ?;