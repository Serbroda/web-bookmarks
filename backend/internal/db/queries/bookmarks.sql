-- name: FindBookmarkById :one
SELECT *
FROM bookmarks b
WHERE id = ? LIMIT 1;

