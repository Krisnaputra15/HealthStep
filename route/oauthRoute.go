package route

import (
	"github.com/Krisnaputra15/gsc-solution/controller"
	"github.com/labstack/echo/v4"
)

func SetGoogleOauthRoutes(e *echo.Echo) {
	e.GET("/auth/google/login", controller.GoogleLogin)
	e.GET("/auth/google/callback", controller.GoogleCallback)
}