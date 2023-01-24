package ohagidinnerprivate

import (
	"net/http"

	"github.com/haton14/ohagi-dinner-private/ohagi-dinner-api/gen/sqlc"
	"github.com/labstack/echo/v4"
)

type menuGetRequest struct {
	DinnerID int64 `param:"id"`
}

type menuGetResponse struct {
	DinnerID  int64        `json:"dinnerId"`
	Menus     []menuDetail `json:"menu"`
	CreatedAt int64        `json:"createdAt"`
	UpdatedAt int64        `json:"updatedAt"`
}

type menuDetail struct {
	FoodID   int64   `json:"foodId"`
	Name     string  `json:"name"`
	Quantity float64 `json:"quantity"`
	Unit     string  `json:"unit"`
}

func (a App) getMenu(c echo.Context) error {
	menuReqeust := menuGetRequest{}
	if err := c.Bind(&menuReqeust); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	dinner, err := a.query.GetDinner(c.Request().Context(), menuReqeust.DinnerID)
	if err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}

	menus, err := a.query.GetMenu(c.Request().Context(), menuReqeust.DinnerID)
	if err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}

	menuDetails := make([]menuDetail, 0, len(menus))
	for _, v := range menus {
		menuDetails = append(menuDetails, menuDetail{
			FoodID:   v.ID,
			Name:     v.Name,
			Quantity: v.Quantity,
			Unit:     v.Unit,
		})
	}

	return c.JSON(http.StatusOK, menuGetResponse{
		DinnerID:  dinner.ID,
		Menus:     menuDetails,
		CreatedAt: dinner.CreatedAt.UnixMicro(),
		UpdatedAt: dinner.UpdatedAt.UnixMicro(),
	})
}

type menuPost struct {
	DinnerID int64   `json:"dinnerId"`
	FoodID   int64   `json:"foodId"`
	Quantity float64 `json:"quantity"`
}

func (a App) createMenu(c echo.Context) error {
	menu := menuPost{}
	if err := c.Bind(&menu); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	err := a.query.CreateMenu(
		c.Request().Context(),
		sqlc.CreateMenuParams{
			DinnerID: menu.DinnerID,
			FoodID:   menu.FoodID,
			Quantity: menu.Quantity,
		},
	)
	if err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}

	return c.NoContent(http.StatusCreated)
}

type menuPatch struct {
	DinnerID int64   `json:"dinnerId"`
	FoodID   int64   `json:"foodId"`
	Quantity float64 `json:"quantity"`
}

func (a App) updateMenu(c echo.Context) error {
	menu := menuPatch{}
	if err := c.Bind(&menu); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	err := a.query.UpdateMenu(
		c.Request().Context(),
		sqlc.UpdateMenuParams{
			DinnerID: menu.DinnerID,
			FoodID:   menu.FoodID,
			Quantity: menu.Quantity,
		},
	)
	if err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}

	return c.NoContent(http.StatusNoContent)
}
