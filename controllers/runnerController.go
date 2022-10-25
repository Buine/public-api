package controllers

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/tesis/user-ms/services"
	"github.com/tesis/user-ms/utils/jwt"
	"net/http"
)

func RunQuery(context echo.Context) error {
	var request interface{}
	err := json.NewDecoder(context.Request().Body).Decode(&request)
	if err != nil {
		return err
	}
	_, err = jwt.ValidateJwt(context)
	if err != nil {
		return err
	}
	query, err := services.RunQuery(request, context.Request().Header)
	if err != nil {
		return err
	}

	context.JSON(http.StatusOK, query)

	return nil
}
