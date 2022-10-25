package controllers

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/tesis/user-ms/errors"
	"github.com/tesis/user-ms/requests"
	"github.com/tesis/user-ms/services"
	"github.com/tesis/user-ms/utils/jwt"
	"net/http"
)

func CreateIntegration(context echo.Context) error {
	var request requests.CreateIntegrationRequest
	err := json.NewDecoder(context.Request().Body).Decode(&request)
	if err != nil {
		return err
	}
	_, err = jwt.ValidateJwt(context)
	if err != nil {
		return err
	}
	integration, err := services.CreateIntegration(request, context.Request().Header)
	if err != nil {
		return err
	}

	context.JSON(http.StatusCreated, integration)

	return nil
}

func GetIntegrations(context echo.Context) error {
	_, err := jwt.ValidateJwt(context)
	if err != nil {
		return err
	}
	integrations, err := services.GetIntegrations(context.Request().Header)
	if err != nil {
		return err
	}

	context.JSON(http.StatusOK, integrations)

	return nil
}

func GetIntegrationByCode(context echo.Context) error {
	code := context.Param("code")
	if code == "" {
		return &errors.BadRequest
	}
	_, err := jwt.ValidateJwt(context)
	if err != nil {
		return err
	}
	integration, err := services.GetIntegrationByCode(code, context.Request().Header)
	if err != nil {
		return err
	}

	context.JSON(http.StatusOK, integration)

	return nil
}

func GetSchemaByCode(context echo.Context) error {
	code := context.Param("code")
	if code == "" {
		return &errors.BadRequest
	}
	_, err := jwt.ValidateJwt(context)
	if err != nil {
		return err
	}
	integration, err := services.GetSchemaByCode(code, context.Request().Header)
	if err != nil {
		return err
	}

	context.JSON(http.StatusOK, integration)

	return nil
}

func UpdateIntegration(context echo.Context) error {
	var request requests.UpdateIntegrationRequest
	err := json.NewDecoder(context.Request().Body).Decode(&request)
	if err != nil {
		return err
	}

	code := context.Param("code")
	if code == "" {
		return &errors.BadRequest
	}

	_, err = jwt.ValidateJwt(context)
	if err != nil {
		return err
	}
	integration, err := services.UpdateIntegration(code, request, context.Request().Header)
	if err != nil {
		return err
	}

	context.JSON(http.StatusOK, integration)

	return nil
}

func DeleteIntegration(context echo.Context) error {
	code := context.Param("code")
	if code == "" {
		return &errors.BadRequest
	}

	_, err := jwt.ValidateJwt(context)
	if err != nil {
		return err
	}
	res, err := services.DeleteIntegration(code, context.Request().Header)
	if err != nil {
		return err
	}

	context.JSON(http.StatusOK, res)

	return nil

}
