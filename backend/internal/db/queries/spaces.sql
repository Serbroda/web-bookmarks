-- name: CreateSpace :one
INSERT INTO spaces(created_at,
                   updated_at,
                   name,
                   description,
                   visibility)
VALUES (CURRENT_TIMESTAMP,
        CURRENT_TIMESTAMP,
        ?,
        ?,
        ?) RETURNING *
;

-- name: CreateSpaceUser :one
INSERT INTO spaces_users(created_at,
                         space_id,
                         user_id,
                         role)
VALUES (CURRENT_TIMESTAMP,
        ?,
        ?,
        ?) RETURNING *
;

-- name: FindSpaceById :one
SELECT *
FROM spaces s
WHERE id = ? LIMIT 1;

-- name: FindSpacesByUserId :many
SELECT s.*--, su.user_id, su.role, su.created_at as assigned_at
FROM spaces s
         INNER JOIN spaces_users su on
    su.space_id = s.id AND
    su.user_id = ?
;