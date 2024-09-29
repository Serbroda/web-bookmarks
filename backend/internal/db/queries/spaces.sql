-- name: FindSpaceById :one
SELECT *
FROM spaces u
WHERE id = ? LIMIT 1;
