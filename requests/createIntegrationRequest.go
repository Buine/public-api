package requests

type CreateIntegrationRequest struct {
	Name     string   `json:"name"`
	Database Database `json:"database"`
}

type Database struct {
	Host     string `json:"host"`
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Ssl      bool   `json:"ssl"`
	Port     int    `json:"port"`
}
