-- name: ListMenu :many
SELECT dinner.id,
  dinner.created_at,
  food.name,
  dinner_menu.quantity,
  food.unit
FROM dinner as dinner
  INNER JOIN dinner_menu ON dinner.id = dinner_menu.dinner_id
  INNER JOIN food ON dinner_menu.food_id = food.id;
-- name: GetMenu :many
SELECT food.id,
  food.name,
  dinner_menu.quantity,
  food.unit
FROM dinner as dinner
  INNER JOIN dinner_menu ON dinner.id = dinner_menu.dinner_id
  INNER JOIN food ON dinner_menu.food_id = food.id
WHERE dinner.id = ?1;
-- name: CreateMenu :exec
INSERT INTO dinner_menu(dinner_id, food_id, quantity)
VALUES(?1, ?2, ?3);
-- name: UpdateMenu :exec
UPDATE dinner_menu
SET quantity = ?1
WHERE dinner_id = ?2
  AND food_id = ?3;
