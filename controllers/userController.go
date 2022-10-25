package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/tesis/user-ms/requests"
	"github.com/tesis/user-ms/services"
	"github.com/tesis/user-ms/utils/jwt"
	"net/http"
)

func CreateUser(context echo.Context) error {
	var userRequest requests.UserRequest
	err := json.NewDecoder(context.Request().Body).Decode(&userRequest)
	if err != nil {
		return err
	}
	user, err := services.CreateUser(userRequest)
	if err != nil {
		return err
	}

	token, err := jwt.BuildToken(*user)
	if err != nil {
		return err
	}
	context.Response().Header().Set("Authorization", fmt.Sprintf("Bearer %s", token))

	context.JSON(http.StatusCreated, user)

	return nil
}

func GetUser(context echo.Context) error {
	token, err := jwt.ValidateJwt(context)
	if err != nil {
		return err
	}
	user, err := services.GetUser(token.Code)
	if err != nil {
		return err
	}

	context.JSON(http.StatusOK, user)

	return nil
}

func Login(context echo.Context) error {
	var userLoginRequest requests.UserLoginRequest
	err := json.NewDecoder(context.Request().Body).Decode(&userLoginRequest)
	if err != nil {
		return err
	}
	user, err := services.GetUserLogin(userLoginRequest)
	if err != nil {
		return err
	}

	token, err := jwt.BuildToken(*user)
	if err != nil {
		return err
	}
	context.Response().Header().Set("Authorization", fmt.Sprintf("Bearer %s", token))

	context.JSON(http.StatusOK, user)

	return nil
}
