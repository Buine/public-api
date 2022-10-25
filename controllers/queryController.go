package controllers

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/tesis/user-ms/errors"
	"github.com/tesis/user-ms/services"
	"github.com/tesis/user-ms/utils/jwt"
	"net/http"
)

func CreateQuery(context echo.Context) error {
	var request interface{}
	err := json.NewDecoder(context.Request().Body).Decode(&request)
	if err != nil {
		return err
	}
	_, err = jwt.ValidateJwt(context)
	if err != nil {
		return err
	}
	query, err := services.CreateQuery(request, context.Request().Header)
	if err != nil {
		return err
	}

	context.JSON(http.StatusOK, query)

	return nil
}

func ListQueriesByIntegration(context echo.Context) error {
	integrationCode := context.Param("integrationCode")
	err := validateCodeFromString(integrationCode)
	if err != nil {
		return err
	}
	_, err = jwt.ValidateJwt(context)
	if err != nil {
		return err
	}
	query, err := services.ListQueriesByIntegration(integrationCode, context.Request().Header)
	if err != nil {
		return err
	}

	context.JSON(http.StatusOK, query)

	return nil
}

func GetQuery(context echo.Context) error {
	integrationCode := context.Param("integrationCode")
	queryCode := context.Param("queryCode")
	withJson := validateBooleanFromString(context.QueryParam("with_json"))
	err := validateCodeFromString(integrationCode)
	if err != nil {
		return err
	}
	err = validateCodeFromString(queryCode)
	if err != nil {
		return err
	}
	_, err = jwt.ValidateJwt(context)
	if err != nil {
		return err
	}
	query, err := services.GetQuery(integrationCode, queryCode, withJson, context.Request().Header)
	if err != nil {
		return err
	}

	context.JSON(http.StatusOK, query)

	return nil
}

func UpdateQuery(context echo.Context) error {
	integrationCode := context.Param("integrationCode")
	queryCode := context.Param("queryCode")
	err := validateCodeFromString(integrationCode)
	if err != nil {
		return err
	}
	err = validateCodeFromString(queryCode)
	if err != nil {
		return err
	}
	var request interface{}
	err = json.NewDecoder(context.Request().Body).Decode(&request)
	if err != nil {
		return err
	}
	_, err = jwt.ValidateJwt(context)
	if err != nil {
		return err
	}
	query, err := services.UpdateQuery(request, integrationCode, queryCode, context.Request().Header)
	if err != nil {
		return err
	}

	context.JSON(http.StatusOK, query)

	return nil
}

func DeleteQuery(context echo.Context) error {
	integrationCode := context.Param("integrationCode")
	queryCode := context.Param("queryCode")
	err := validateCodeFromString(integrationCode)
	if err != nil {
		return err
	}
	err = validateCodeFromString(queryCode)
	if err != nil {
		return err
	}
	_, err = jwt.ValidateJwt(context)
	if err != nil {
		return err
	}
	query, err := services.DeleteQuery(integrationCode, queryCode, context.Request().Header)
	if err != nil {
		return err
	}

	context.JSON(http.StatusOK, query)

	return nil
}

func validateCodeFromString(code string) error {
	if code == "" {
		return &errors.BadRequest
	}
	return nil
}

func validateBooleanFromString(boolean string) bool {
	return boolean == "true"
}
