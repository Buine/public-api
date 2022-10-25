package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tesis/user-ms/responses"
)

func HealthCheck(context echo.Context) error {
	context.JSON(http.StatusOK, responses.HealthResponse{
		Status: "Ok",
	})
	return nil
}