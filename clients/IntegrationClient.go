package clients

import (
	"fmt"
	"github.com/gojek/heimdall/v7/httpclient"
	"github.com/tesis/user-ms/constants/routes"
	"github.com/tesis/user-ms/requests"
	"github.com/tesis/user-ms/responses"
	"net/http"
)

func CreateIntegration(integrationRequest requests.CreateIntegrationRequest, headers http.Header) (*responses.CreateIntegrationResponse, error) {
	var response responses.CreateIntegrationResponse
	body, err := InterfaceToReader(integrationRequest)
	if err != nil {
		return nil, err
	}
	client := httpclient.NewClient()
	res, err := client.Post(fmt.Sprintf(routes.CreateIntegration, ExternalMs.IntegrationMs), body, headers)
	if err != nil {
		return nil, err
	}
	err = ValidateHttpResponse(res, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func GetIntegrations(headers http.Header) (*responses.ListIntegrationResponse, error) {
	var response responses.ListIntegrationResponse
	client := httpclient.NewClient()
	res, err := client.Get(fmt.Sprintf(routes.GetIntegrations, ExternalMs.IntegrationMs), headers)
	if err != nil {
		return nil, err
	}
	err = ValidateHttpResponse(res, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func GetIntegrationByCode(code string, headers http.Header) (*responses.IntegrationResponse, error) {
	var response responses.IntegrationResponse
	client := httpclient.NewClient()
	res, err := client.Get(fmt.Sprintf(routes.GetIntegrationByCode, ExternalMs.IntegrationMs, code), headers)
	if err != nil {
		return nil, err
	}
	err = ValidateHttpResponse(res, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func GetSchemaByCode(code string, headers http.Header) (interface{}, error) {
	var response interface{}
	client := httpclient.NewClient()
	res, err := client.Get(fmt.Sprintf(routes.GetSchemaByCode, ExternalMs.IntegrationMs, code), headers)
	if err != nil {
		return nil, err
	}
	err = ValidateHttpResponse(res, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func UpdateIntegration(code string, updateIntegration requests.UpdateIntegrationRequest, headers http.Header) (*responses.IntegrationResponse, error) {
	var response responses.IntegrationResponse
	body, err := InterfaceToReader(updateIntegration)
	if err != nil {
		return nil, err
	}
	client := httpclient.NewClient()
	res, err := client.Patch(fmt.Sprintf(routes.GetIntegrationByCode, ExternalMs.IntegrationMs, code), body, headers)
	if err != nil {
		return nil, err
	}
	err = ValidateHttpResponse(res, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func DeleteIntegration(code string, headers http.Header) (*responses.IntegrationResponse, error) {
	var response responses.IntegrationResponse
	client := httpclient.NewClient()
	res, err := client.Delete(fmt.Sprintf(routes.GetIntegrationByCode, ExternalMs.IntegrationMs, code), headers)
	if err != nil {
		return nil, err
	}
	err = ValidateHttpResponse(res, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
