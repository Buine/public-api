package responses

import "time"

type IntegrationResponse struct {
	Code      string    `json:"code"`
	Name      string    `json:"name"`
	Status    string    `json:"status"`
	Checksum  string    `json:"checksum"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Config    Config    `json:"config"`
}
type Config struct {
	Host     string `json:"host"`
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Ssl      bool   `json:"ssl"`
	Port     int    `json:"port"`
}
