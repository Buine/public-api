package routes

import "fmt"

const HEALTH_CHECK = "/health-check"

const V1 = "/v1"

const USER = "/user"

var LOGIN = fmt.Sprint(USER, "/login")

const GET_USER = "/user"

const INTEGRATION = "/integration"
const INTEGRATION_BY_CODE = "/integration/:code"
const SCHEMA_BY_INTEGRATION_CODE = "/integration/:code/schema"

const RUN_QUERY = "/query/run"
const QUERY = "/query"
const QUERY_BY_INTEGRATION = "/integration/:integrationCode/queries"
const QUERY_BY_INTEGRATION_CODE_AND_QUERY_CODE = "/integration/:integrationCode/query/:queryCode"
