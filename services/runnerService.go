package services

import (
	"github.com/tesis/user-ms/clients"
	"net/http"
)

func RunQuery(queryRequest interface{}, headers http.Header) (interface{}, error) {
	res, err := clients.RunQuery(queryRequest, headers)
	if err != nil {
		return nil, err
	}
	return res, nil
}
