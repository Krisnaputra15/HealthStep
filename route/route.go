package route

import "github.com/labstack/echo/v4"

func InitRoute(e *echo.Echo) {
	SetGoogleOauthRoutes(e)
}