package services

import (
	"github.com/tesis/user-ms/clients"
	"net/http"
)

func CreateQuery(queryCreateRequest interface{}, headers http.Header) (interface{}, error) {
	res, err := clients.CreateQuery(queryCreateRequest, headers)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func ListQueriesByIntegration(integrationCode string, headers http.Header) (interface{}, error) {
	res, err := clients.ListQueriesByIntegration(integrationCode, headers)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func GetQuery(integrationCode string, queryCode string, withJson bool, headers http.Header) (interface{}, error) {
	res, err := clients.GetQuery(integrationCode, queryCode, withJson, headers)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func UpdateQuery(request interface{}, integrationCode string, queryCode string, headers http.Header) (interface{}, error) {
	res, err := clients.UpdateQuery(request, integrationCode, queryCode, headers)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func DeleteQuery(integrationCode string, queryCode string, headers http.Header) (interface{}, error) {
	res, err := clients.DeleteQuery(integrationCode, queryCode, headers)
	if err != nil {
		return nil, err
	}
	return res, nil
}
