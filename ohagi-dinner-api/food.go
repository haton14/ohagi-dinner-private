package ohagidinnerprivate

import (
	"net/http"
	"sort"

	"github.com/haton14/ohagi-dinner-private/ohagi-dinner-api/gen/sqlc"
	"github.com/haton14/ohagi-dinner-private/ohagi-dinner-api/utility"
	"github.com/labstack/echo/v4"
)

type foodList struct {
	Foods []food `json:"foods"`
}

type food struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Unit      string `json:"unit"`
	CreatedAt int64  `json:"createdAt"`
	UpdatedAt int64  `json:"updatedAt"`
}

func (a App) listFood(c echo.Context) error {
	data, err := a.query.ListFood(c.Request().Context())
	if err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}

	sort.Slice(data, func(i, j int) bool { return data[i].ID < data[j].ID })

	foods := make([]food, 0, len(data))
	for _, v := range data {
		foods = append(foods, food{
			ID:        v.ID,
			Name:      v.Name,
			Unit:      v.Unit,
			CreatedAt: v.CreatedAt.UnixMicro(),
			UpdatedAt: v.CreatedAt.UnixMicro(),
		})
	}
	return c.JSON(http.StatusOK, foods)
}

type foodPost struct {
	Name string `json:"name"`
	Unit string `json:"unit"`
}

func (a App) createFood(c echo.Context) error {
	food := foodPost{}
	if err := c.Bind(&food); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	now := utility.NowInJST()
	err := a.query.CreateFood(
		c.Request().Context(),
		sqlc.CreateFoodParams{
			Name:      food.Name,
			Unit:      food.Unit,
			CreatedAt: now,
			UpdatedAt: now,
		},
	)
	if err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}

	return c.NoContent(http.StatusCreated)
}
