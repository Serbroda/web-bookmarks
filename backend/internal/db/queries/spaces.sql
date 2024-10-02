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
FROM spaces u
WHERE id = ? LIMIT 1;

