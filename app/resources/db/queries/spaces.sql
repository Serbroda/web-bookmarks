-- name: FindSpace :one
SELECT *
FROM spaces
WHERE id = ?
LIMIT 1;
-- name: FindSpaceByShortId :one
SELECT *
FROM spaces
WHERE short_id = ?
LIMIT 1;
-- name: FindUserSpaces :many
SELECT s.*
FROM spaces s
    INNER JOIN users_spaces us on us.space_id = s.id
WHERE us.user_id = ?;
-- name: FindUserSpace :one
SELECT s.*
FROM spaces s
    INNER JOIN users_spaces us on us.space_id = s.id
WHERE us.user_id = ?
    and s.short_id = ?
    and (
        s.visibility <> 'PRIVATE'
        or s.owner_id = us.user_id
    )
LIMIT 1;
-- name: CreateSpace :execlastid
INSERT INTO spaces (
        created_at,
        short_id,
        owner_id,
        name,
        description,
        visibility
    )
VALUES(CURRENT_TIMESTAMP, ?, ?, ?, ?, ?);
-- name: UpdateSpace :exec
UPDATE spaces
SET updated_at = CURRENT_TIMESTAMP,
    name = COALESCE(?, updated_at),
    description = COALESCE(?, updated_at),
    visibility = COALESCE(?, updated_at)
WHERE id = ?;
-- name: InsertUserSpace :exec
INSERT INTO users_spaces (user_id, space_id, role_id, created_at)
VALUES(?, ?, ?, CURRENT_TIMESTAMP);
-- name: DeleteUserSpace :exec
DELETE FROM users_spaces
WHERE user_id = ?
    AND space_id = ?;