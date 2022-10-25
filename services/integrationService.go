package services

import (
	"github.com/tesis/user-ms/clients"
	"github.com/tesis/user-ms/requests"
	"github.com/tesis/user-ms/responses"
	"net/http"
)

func CreateIntegration(userRequest requests.CreateIntegrationRequest, headers http.Header) (*responses.CreateIntegrationResponse, error) {
	res, err := clients.CreateIntegration(userRequest, headers)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func GetIntegrations(headers http.Header) (*responses.ListIntegrationResponse, error) {
	res, err := clients.GetIntegrations(headers)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func GetIntegrationByCode(code string, headers http.Header) (*responses.IntegrationResponse, error) {
	res, err := clients.GetIntegrationByCode(code, headers)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func GetSchemaByCode(code string, headers http.Header) (interface{}, error) {
	res, err := clients.GetSchemaByCode(code, headers)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func UpdateIntegration(code string, updateIntegration requests.UpdateIntegrationRequest, headers http.Header) (*responses.IntegrationResponse, error) {
	res, err := clients.UpdateIntegration(code, updateIntegration, headers)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func DeleteIntegration(code string, headers http.Header) (*responses.IntegrationResponse, error) {
	res, err := clients.DeleteIntegration(code, headers)
	if err != nil {
		return nil, err
	}
	return res, nil
}
