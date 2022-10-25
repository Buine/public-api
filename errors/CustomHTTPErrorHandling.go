package errors

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func CustomHTTPErrorHandler(err error, context echo.Context) {
	if context.Response().Committed {
		return
	}
	code := http.StatusInternalServerError
	if customError, ok := err.(*ErrorResponse); ok {
		fmt.Println(customError)
		statusCode := customError.Status
		if customError.Status == 0 {
			statusCode = 400
		}
		context.JSON(statusCode, customError)
		return
	}
	fmt.Println(err)

	if errorEcho, ok := err.(*echo.HTTPError); ok {
		errorResponse := &ErrorResponse{
			Code:     fmt.Sprint(errorEcho.Code),
			Messages: []string{fmt.Sprintf("%+v", errorEcho.Message)},
		}
		context.JSON(errorEcho.Code, errorResponse)
		return
	}

	context.JSON(code, &ErrorResponse{
		Code:     "INTERNAL_ERROR",
		Messages: []string{fmt.Sprintf("%+v", err)},
	})
}
