-- name: ListMenu :many
SELECT
  dinner.id,
  dinner.created_at,
  food.name,
  dinner_menu.quantity,
  food.unit
FROM dinner as dinner
INNER JOIN dinner_menu ON dinner.id = dinner_menu.dinner_id
INNER JOIN food ON dinner_menu.food_id = food.id;
