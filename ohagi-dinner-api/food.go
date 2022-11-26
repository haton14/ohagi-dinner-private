package ohagidinnerprivate

import (
	"net/http"

	"github.com/haton14/ohagi-dinner-private/ohagi-dinner-api/gen/sqlc"
	"github.com/haton14/ohagi-dinner-private/ohagi-dinner-api/utility"
	"github.com/labstack/echo/v4"
)

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
