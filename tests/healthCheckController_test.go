package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tesis/user-ms/utils/file"

	"github.com/labstack/echo/v4"
	"github.com/tesis/user-ms/constants/routes"
	"github.com/tesis/user-ms/controllers"
	"github.com/tesis/user-ms/responses"
)

func TestHealthCheck(t *testing.T) {
	healthCheckJSON, _ := file.ReadFile("./resources/response/health_check.json")
	echoTest := echo.New()
	req := httptest.NewRequest(http.MethodGet, routes.HEALTH_CHECK, nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	context := echoTest.NewContext(req, res)

	if assert.NoError(t, controllers.HealthCheck(context)) {
		assert.Equal(t, http.StatusOK, res.Code)

		// Marshall
		responseExpect := new(responses.HealthResponse)
		responseActual := new(responses.HealthResponse)
		json.Unmarshal([]byte(healthCheckJSON), &responseExpect)
		json.Unmarshal(res.Body.Bytes(), &responseActual)

		assert.Equal(t, responseExpect, responseActual)

	}
}
