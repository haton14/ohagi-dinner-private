-- name: CreateDinnerAndReturnID :one
INSERT INTO dinner (created_at, updated_at)
VALUES(?1, ?2)
RETURNING id;
-- name: GetDinner :one
SELECT *
FROM dinner
WHERE id = ?1;
