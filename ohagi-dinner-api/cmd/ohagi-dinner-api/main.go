package main

import (
	"database/sql"
	"log"

	ohagidinnerprivate "github.com/haton14/ohagi-dinner-private/ohagi-dinner-api"
	"github.com/haton14/ohagi-dinner-private/ohagi-dinner-api/gen/sqlc"
	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./dinner.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	e := echo.New()
	app := ohagidinnerprivate.NewApp(e, sqlc.New(db))
	app.Start()
}
