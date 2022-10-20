-- name: CreateFood :exec
INSERT INTO food (name, unit, created_at, updated_at)
VALUES(?, ?, ?, ?);
-- name: ListFood :many
SELECT *
FROM food;
