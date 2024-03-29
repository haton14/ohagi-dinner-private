// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: menu.sql

package sqlc

import (
	"context"
	"time"
)

const createMenu = `-- name: CreateMenu :exec
INSERT INTO dinner_menu(dinner_id, food_id, quantity)
VALUES(?1, ?2, ?3)
`

type CreateMenuParams struct {
	DinnerID int64
	FoodID   int64
	Quantity float64
}

func (q *Queries) CreateMenu(ctx context.Context, arg CreateMenuParams) error {
	_, err := q.db.ExecContext(ctx, createMenu, arg.DinnerID, arg.FoodID, arg.Quantity)
	return err
}

const getMenu = `-- name: GetMenu :many
SELECT food.id,
  food.name,
  dinner_menu.quantity,
  food.unit
FROM dinner as dinner
  INNER JOIN dinner_menu ON dinner.id = dinner_menu.dinner_id
  INNER JOIN food ON dinner_menu.food_id = food.id
WHERE dinner.id = ?1
`

type GetMenuRow struct {
	ID       int64
	Name     string
	Quantity float64
	Unit     string
}

func (q *Queries) GetMenu(ctx context.Context, id int64) ([]GetMenuRow, error) {
	rows, err := q.db.QueryContext(ctx, getMenu, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetMenuRow
	for rows.Next() {
		var i GetMenuRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Quantity,
			&i.Unit,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listMenu = `-- name: ListMenu :many
SELECT dinner.id,
  dinner.created_at,
  food.name,
  dinner_menu.quantity,
  food.unit
FROM dinner as dinner
  LEFT JOIN dinner_menu ON dinner.id = dinner_menu.dinner_id
  LEFT JOIN food ON dinner_menu.food_id = food.id
`
// sqlcのバグがあるので一時的に手動修正
// https://github.com/kyleconroy/sqlc/issues/2044
type ListMenuRow struct {
	ID        int64
	CreatedAt time.Time
	Name      *string
	Quantity  *float64
	Unit      *string
}

func (q *Queries) ListMenu(ctx context.Context) ([]ListMenuRow, error) {
	rows, err := q.db.QueryContext(ctx, listMenu)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListMenuRow
	for rows.Next() {
		var i ListMenuRow
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.Name,
			&i.Quantity,
			&i.Unit,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateMenu = `-- name: UpdateMenu :exec
UPDATE dinner_menu
SET quantity = ?1
WHERE dinner_id = ?2
  AND food_id = ?3
`

type UpdateMenuParams struct {
	Quantity float64
	DinnerID int64
	FoodID   int64
}

func (q *Queries) UpdateMenu(ctx context.Context, arg UpdateMenuParams) error {
	_, err := q.db.ExecContext(ctx, updateMenu, arg.Quantity, arg.DinnerID, arg.FoodID)
	return err
}
