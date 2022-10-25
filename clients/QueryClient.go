package clients

import (
	"fmt"
	"github.com/gojek/heimdall/v7/httpclient"
	"github.com/tesis/user-ms/constants/routes"
	"net/http"
)

func CreateQuery(queryCreateRequest interface{}, headers http.Header) (interface{}, error) {
	var response interface{}
	body, err := InterfaceToReader(queryCreateRequest)
	if err != nil {
		return nil, err
	}
	client := httpclient.NewClient()
	res, err := client.Post(fmt.Sprintf(routes.CreateQuery, ExternalMs.QueryMs), body, headers)
	if err != nil {
		return nil, err
	}
	err = ValidateHttpResponse(res, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func ListQueriesByIntegration(integrationCode string, headers http.Header) (interface{}, error) {
	var response interface{}
	client := httpclient.NewClient()
	res, err := client.Get(fmt.Sprintf(routes.ListQueriesByIntegrationCode, ExternalMs.QueryMs, integrationCode), headers)
	if err != nil {
		return nil, err
	}
	err = ValidateHttpResponse(res, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func GetQuery(integrationCode string, queryCode string, withJson bool, headers http.Header) (interface{}, error) {
	var response interface{}
	client := httpclient.NewClient()
	url := fmt.Sprint(fmt.Sprintf(routes.QueryByCodeAndIntegrationCode, ExternalMs.QueryMs, queryCode, integrationCode), "?with_json=", withJson)
	res, err := client.Get(url, headers)
	if err != nil {
		return nil, err
	}
	err = ValidateHttpResponse(res, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func UpdateQuery(request interface{}, integrationCode string, queryCode string, headers http.Header) (interface{}, error) {
	var response interface{}
	body, err := InterfaceToReader(request)
	if err != nil {
		return nil, err
	}
	client := httpclient.NewClient()
	res, err := client.Patch(fmt.Sprintf(routes.QueryByCodeAndIntegrationCode, ExternalMs.QueryMs, queryCode, integrationCode), body, headers)
	if err != nil {
		return nil, err
	}
	err = ValidateHttpResponse(res, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func DeleteQuery(integrationCode string, queryCode string, headers http.Header) (interface{}, error) {
	var response interface{}
	client := httpclient.NewClient()
	res, err := client.Delete(fmt.Sprintf(routes.QueryByCodeAndIntegrationCode, ExternalMs.QueryMs, queryCode, integrationCode), headers)
	if err != nil {
		return nil, err
	}
	err = ValidateHttpResponse(res, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
