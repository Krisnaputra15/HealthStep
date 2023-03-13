package main

import (
	"HealthStep/config"
	"HealthStep/db"
	"HealthStep/migration"
	"HealthStep/route"

	"github.com/labstack/echo/v4"
)

func main() {
	config.LoadEnv()
	config.LoadOauth2Config()
	db.ConnectToDB()
	migration.Migrate()

	e := echo.New()

	route.InitRoute(e)

	e.Logger.Fatal(e.Start(":8000"))
}
