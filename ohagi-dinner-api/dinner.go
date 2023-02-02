package ohagidinnerprivate

import (
	"fmt"
	"net/http"
	"sort"
	"time"

	"github.com/haton14/ohagi-dinner-private/ohagi-dinner-api/gen/sqlc"
	"github.com/haton14/ohagi-dinner-private/ohagi-dinner-api/utility"
	"github.com/labstack/echo/v4"
)

type dinnerList struct {
	Dinners []dinner `json:"dinners"`
}

type dinner struct {
	ID        int64  `json:"id"`
	Menus     []menu `json:"menus"`
	CreatedAt int64  `json:"createdAt"`
}

type menu struct {
	Name     string  `json:"name"`
	Quantity float64 `json:"quantity"`
	Unit     string  `json:"unit"`
}

func (a App) listDinner(c echo.Context) error {
	data, err := a.query.ListMenu(c.Request().Context())
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "internal server error\n")
	}

	sort.SliceStable(data, func(i, j int) bool {
		if data[i].Name == nil {
			return true
		}
		if data[j].Name == nil {
			return false
		}
		return *data[i].Name < *data[j].Name
	})
	sort.SliceStable(data, func(i, j int) bool { return data[i].ID < data[j].ID })

	// menuレコードの1/3もdinnerはないので最低限の容量を事前確保する
	dinners := make([]dinner, 0, len(data)/3)
	var currentDinnerID int64
	for _, v := range data {
		if currentDinnerID != v.ID {
			dinners = append(dinners, dinner{
				ID:        v.ID,
				Menus:     make([]menu, 0, 3),
				CreatedAt: v.CreatedAt.Unix(),
			})
			currentDinnerID = v.ID
		}
		if v.Name != nil {
			dinners[len(dinners)-1].Menus = append(dinners[len(dinners)-1].Menus, menu{
				Name:     *v.Name,
				Quantity: *v.Quantity,
				Unit:     *v.Unit,
			})
		}
	}
	return c.JSON(http.StatusOK, dinnerList{Dinners: dinners})
}

type dinnerPost struct {
	CreatedAt *int64 `json:"createdAt"`
}

func (a App) createDinner(c echo.Context) error {
	dinner := dinnerPost{}
	if err := c.Bind(&dinner); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	now := time.Now()
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
		"dinnerId": dinnerID,
	})
}

type dinnerGet struct {
	ID        int64 `param:"id" json:"id"`
	CreatedAt int64 `json:"createdAt"`
	UpdatedAt int64 `json:"updatedAt"`
}

func (a App) getDinner(c echo.Context) error {
	dinner := dinnerGet{}
	if err := c.Bind(&dinner); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	d, err := a.query.GetDinner(
		c.Request().Context(),
		dinner.ID,
	)
	if err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	dinner.CreatedAt = d.CreatedAt.UnixMicro()
	dinner.UpdatedAt = d.UpdatedAt.UnixMicro()
	return c.JSON(http.StatusCreated, dinner)
}
