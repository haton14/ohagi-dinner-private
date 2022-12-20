package ohagidinnerprivate

import (
	"fmt"
	"log"
	"os"

	"github.com/haton14/ohagi-dinner-private/ohagi-dinner-api/gen/sqlc"
	"github.com/labstack/echo/v4"
)

type App struct {
	echo  *echo.Echo
	query *sqlc.Queries
}

func NewApp(echo *echo.Echo, query *sqlc.Queries) App {
	return App{echo: echo, query: query}
}

func (a App) Start() {
	a.echo.GET("/dinners", a.listDinner)
	a.echo.POST("/dinners", a.createDinner)

	a.echo.GET("/foods", a.listFood)
	a.echo.POST("/foods", a.createFood)

	a.echo.GET("/menus/:id", a.getMenu)
	a.echo.POST("/menus", a.createMenu)
	a.echo.PATCH("/menus", a.updateMenu)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Panicln(a.echo.Start(fmt.Sprint(":", port)))
}
