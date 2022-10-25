package requests

type UpdateIntegrationRequest struct {
	Name   *string `json:"name"`
	Config *Config `json:"config"`
}
type Config struct {
	Host     *string `json:"host"`
	Username *string `json:"username"`
	Password *string `json:"password"`
	Name     *string `json:"name"`
	Ssl      *bool   `json:"ssl"`
	Port     *int    `json:"port"`
}
