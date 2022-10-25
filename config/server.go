package config

import (
	"os"
)

type ServerConfig struct {
	MicroserviceName string
	JwtKey           string
	Port             string
}

func (current *ServerConfig) Init() ServerConfig {
	current.MicroserviceName = "public-api"
	if current.Port = os.Getenv("SERVER_PORT"); current.Port == "" {
		panic("Environment Key: SERVER_PORT, Not loaded")
	}
	if current.JwtKey = os.Getenv("JWT_KEY"); current.JwtKey == "" {
		panic("Environment Key: JWT_KEY, Not loaded")
	}

	return *current
}

type ExternalMs struct {
	UserMs        string
	IntegrationMs string
	RunnerMs      string
	QueryMs       string
}

func (current *ExternalMs) Init() ExternalMs {
	if current.UserMs = os.Getenv("USER_MS_URL"); current.UserMs == "" {
		panic("Environment Key: USER_MS_URL, Not loaded")
	}
	if current.IntegrationMs = os.Getenv("INTEGRATION_MS_URL"); current.IntegrationMs == "" {
		panic("Environment Key: INTEGRATION_MS_URL, Not loaded")
	}
	if current.RunnerMs = os.Getenv("RUNNER_MS_URL"); current.RunnerMs == "" {
		panic("Environment Key: RUNNER_MS_URL, Not loaded")
	}
	if current.QueryMs = os.Getenv("QUERY_MS_URL"); current.QueryMs == "" {
		panic("Environment Key: QUERY_MS_URL, Not loaded")
	}

	return *current
}
