package main

import (
	"github.com/Krisnaputra15/gsc-solution/config"
	// "github.com/Krisnaputra15/gsc-solution/db"
	"github.com/Krisnaputra15/gsc-solution/route"
	"github.com/labstack/echo/v4"
)

func main() {
	config.LoadEnv()
	config.LoadOauth2Config()
	// db.ConnectToDB()

	e := echo.New()

	route.InitRoute(e)

	e.Logger.Fatal(e.Start(":8000"))
}