package ohagidinnerprivate

import (
	"net/http"
	"sort"

	"github.com/haton14/ohagi-dinner-private/ohagi-dinner-api/gen/sqlc"
	"github.com/haton14/ohagi-dinner-private/ohagi-dinner-api/utility"
	"github.com/labstack/echo/v4"
)

type dinnerList struct {
	Dinners []dinner `json:"dinners"`
}

type dinner struct {
	ID    int64  `json:"id"`
	Menus []menu `json:"menus"`
}

type menu struct {
	Name     string  `json:"name"`
	Quantity float64 `json:"quantity"`
	Unit     string  `json:"unit"`
}

func (a App) listDinner(c echo.Context) error {
	data, err := a.query.ListMenu(c.Request().Context())
	if err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}

	sort.SliceStable(data, func(i, j int) bool { return data[i].Name < data[j].Name })
	sort.SliceStable(data, func(i, j int) bool { return data[i].ID < data[j].ID })

	// menuレコードの1/3もdinnerはないので最低限の容量を事前確保する
	dinners := make([]dinner, 0, len(data)/3)
	var currentDinnerID int64
	menus := make([]menu, 0, 3)
	for _, v := range data {
		if currentDinnerID != v.ID {
			dinners = append(dinners, dinner{
				ID:    currentDinnerID,
				Menus: menus,
			})
			currentDinnerID = v.ID
			menus = make([]menu, 0, 3)
		}
		menus = append(menus, menu{
			Name:     v.Name,
			Quantity: v.Quantity,
			Unit:     v.Unit,
		})

	}
	return c.JSON(http.StatusOK, dinnerList{Dinners: dinners})
}

type dinnerPost struct {
	CreatedAt *int64 `json:"created_at"`
}

func (a App) createDinner(c echo.Context) error {
	dinner := dinnerPost{}
	if err := c.Bind(&dinner); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	now := utility.NowInJST()
	dinnerID, err := a.query.CreateDinnerAndReturnID(
		c.Request().Context(),
		sqlc.CreateDinnerAndReturnIDParams{
			CreatedAt: utility.CreatedAt(dinner.CreatedAt, now),
			UpdatedAt: now,
		})
	if err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}

	return c.JSON(http.StatusCreated, map[string]any{
		"dinner_id": dinnerID,
	})
}
