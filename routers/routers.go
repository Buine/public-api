package routers

import (
	"github.com/labstack/echo/v4"
	"github.com/tesis/user-ms/config"
	"github.com/tesis/user-ms/constants/routes"
	"github.com/tesis/user-ms/controllers"
)

var serverConfig = new(config.ServerConfig).Init()

func ConfigRouters(server *echo.Echo) {
	group := server.Group(serverConfig.MicroserviceName)
	group.GET(routes.HEALTH_CHECK, controllers.HealthCheck)

	v1 := group.Group(routes.V1)
	v1.POST(routes.USER, controllers.CreateUser)
	v1.GET(routes.GET_USER, controllers.GetUser)
	v1.POST(routes.LOGIN, controllers.Login)

	v1.POST(routes.INTEGRATION, controllers.CreateIntegration)
	v1.GET(routes.INTEGRATION, controllers.GetIntegrations)
	v1.GET(routes.INTEGRATION_BY_CODE, controllers.GetIntegrationByCode)
	v1.GET(routes.SCHEMA_BY_INTEGRATION_CODE, controllers.GetSchemaByCode)
	v1.PATCH(routes.INTEGRATION_BY_CODE, controllers.UpdateIntegration)
	v1.DELETE(routes.INTEGRATION_BY_CODE, controllers.DeleteIntegration)

	v1.POST(routes.RUN_QUERY, controllers.RunQuery)
	v1.POST(routes.QUERY, controllers.CreateQuery)
	v1.GET(routes.QUERY_BY_INTEGRATION, controllers.ListQueriesByIntegration)
	v1.GET(routes.QUERY_BY_INTEGRATION_CODE_AND_QUERY_CODE, controllers.GetQuery)
	v1.PATCH(routes.QUERY_BY_INTEGRATION_CODE_AND_QUERY_CODE, controllers.UpdateQuery)
	v1.DELETE(routes.QUERY_BY_INTEGRATION_CODE_AND_QUERY_CODE, controllers.DeleteQuery)
}
