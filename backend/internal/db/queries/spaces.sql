-- name: FindSpaceById :one
SELECT *
FROM spaces u
WHERE id = ? LIMIT 1;

INSERT INTO spaces(created_at,
                   updated_at,
                   name,
                   description,
                   owner_id,
                   visibility)
VALUES (CURRENT_TIMESTAMP,
        CURRENT_TIMESTAMP,
        ?,
        ?,
        ?,
        ?)
;