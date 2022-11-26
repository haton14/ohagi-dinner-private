// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package sqlc

import (
	"time"
)

type Dinner struct {
	ID        int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

type DinnerMenu struct {
	DinnerID int64
	FoodID   int64
	Quantity float64
}

type Food struct {
	ID        int64
	Name      string
	Unit      string
	CreatedAt time.Time
	UpdatedAt time.Time
}