package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Krisnaputra15/gsc-solution/config"
	"github.com/Krisnaputra15/gsc-solution/entity"
	"github.com/Krisnaputra15/gsc-solution/model"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type UserData struct {
	User        model.APIUser
	UserProfile model.APIUserProfile
}

type UserReturn struct {
	User  UserData
	Token string
}

func GoogleLogin(c echo.Context) error {
	state := GenerateOauthCookie(c)
	url := config.GoogleOauth2Config.AuthCodeURL(state)
	return c.Redirect(http.StatusTemporaryRedirect, url)
}

func GoogleCallback(c echo.Context) error {
	oauthState, _ := c.Cookie("oauthstate")
	if c.FormValue("state") != oauthState.Value {
		log.Println("invalid oauth google state", c.FormValue("state"))
		return c.Redirect(http.StatusTemporaryRedirect, "http://localhost:8000/")
	}

	data, err := GetUserDataFromGoogle(c.FormValue("code"))
	if err != nil {
		log.Println(err.Error())
		return c.Redirect(http.StatusTemporaryRedirect, "http://localhost:8000/")
	}

	// unmarshalling json from google
	var userDetail entity.User
	json.Unmarshal(data, &userDetail)

	userData, err := SignInUser(userDetail)
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.SetErrorResponse(http.StatusBadRequest, "error", err.Error()))
	}

	// unmarshalling json from SignInUser
	var userReturn UserReturn
	json.Unmarshal(userData, &userReturn)

	return c.JSON(http.StatusOK, entity.SetResponse(http.StatusOK, "login success", userDetail))
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

func SignInUser(userDetail entity.User) ([]byte, error) {
	var user model.User
	var userData UserData

	if len(user.Email) == 0 {
		createUser, err := model.UserCreate(userDetail)
		if err != nil {
			return nil, err
		}

		createUserProfile, err := model.UserProfileCreate(userDetail)
		if err != nil {
			return nil, err
		}

		userData.User = createUser
		userData.UserProfile = createUserProfile
	}

	// create claims for jwt from authenticated user
	claims := &jwt.RegisteredClaims{
		ID:        userData.User.ID,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
	}

	// create token with claims
	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// generate encoded token and send it along with response
	token, err := rawToken.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return nil, err
	}

	data := map[string]interface{}{
		"user_data": userData,
		"token":     token,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}
