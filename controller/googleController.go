package controller

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Krisnaputra15/gsc-solution/config"
	"github.com/Krisnaputra15/gsc-solution/entity"
	"github.com/labstack/echo/v4"
)

func GoogleLogin(c echo.Context) error {
	state := GenerateOauthCookie(c)
	url := config.GoogleOauth2Config.AuthCodeURL(state)
	return c.Redirect(http.StatusTemporaryRedirect, url)
}

func GoogleCallback(c echo.Context) error {
	oauthState, _ := c.Cookie("oauthstate")
	if c.FormValue("state") != oauthState.Value {
		log.Println("invalid oauth google state")
		return c.Redirect(http.StatusTemporaryRedirect, "http://localhost:8000/")
	}

	data, err := GetUserDataFromGoogle(c.FormValue("code"))
	if err != nil {
		log.Println(err.Error())
		return c.Redirect(http.StatusTemporaryRedirect, "http://localhost:8000/")
	}

	// unmarshalling json
	// var userDetail entity.User
	// json.Unmarshal(data, &userDetail)

	return c.JSON(http.StatusOK, entity.SetResponse(http.StatusOK, "success retrieving user data", string(data)))
}

func GetUserDataFromGoogle(code string) ([]byte, error) {
	token, err := config.GoogleOauth2Config.Exchange(context.Background(), code)
	if err != nil {
		return nil, fmt.Errorf("code exchange wrong: %s", err.Error())
	}

	response, err := http.Get(config.GoogleOauth2UrlAPI + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("error getting user data: %s", err.Error())
	}

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading user data: %s", err.Error())
	}

	return contents, nil
}
