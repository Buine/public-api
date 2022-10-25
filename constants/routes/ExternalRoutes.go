package routes

const CreateUserRoute = "%s/v1/user"
const LoginUserRoute = "%s/v1/user/login"
const GetUser = "%s/v1/user/%s"

const CreateIntegration = "%s/v1/integration"
const GetIntegrations = "%s/v1/integration"
const GetIntegrationByCode = "%s/v1/integration/%s"
const GetSchemaByCode = "%s/v1/integration/%s/schema"

const RunQuery = "%s/v1/runner/query"
const CreateQuery = "%s/v1/query"
const ListQueriesByIntegrationCode = "%s/v1/query/integration/%s"
const QueryByCodeAndIntegrationCode = "%s/v1/query/%s/integration/%s"
