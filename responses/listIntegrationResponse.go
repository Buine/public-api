package responses

import "time"

type ListIntegrationResponse []struct {
	Code         string    `json:"code"`
	Name         string    `json:"name"`
	Status       string    `json:"status"`
	NameDatabase string    `json:"name_database"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
