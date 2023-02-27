package controller

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func GenerateOauthCookie(c echo.Context) string {
	var expiration = time.Now().Add(7 * time.Hour)

	b := make([]byte,16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)

	cookie := new(http.Cookie)
	cookie.Name = "oauthstate"
	cookie.Value = state
	cookie.Expires = expiration
	c.SetCookie(cookie)

	return state
}

// func GenerateAuthorizationCookie(c echo.Context) error {

// }