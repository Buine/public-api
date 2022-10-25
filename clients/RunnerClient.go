package clients

import (
	"fmt"
	"github.com/gojek/heimdall/v7/httpclient"
	"github.com/tesis/user-ms/constants/routes"
	"net/http"
)

func RunQuery(queryRequest interface{}, headers http.Header) (interface{}, error) {
	var response interface{}
	body, err := InterfaceToReader(queryRequest)
	if err != nil {
		return nil, err
	}
	client := httpclient.NewClient()
	res, err := client.Post(fmt.Sprintf(routes.RunQuery, ExternalMs.RunnerMs), body, headers)
	if err != nil {
		return nil, err
	}
	err = ValidateHttpResponse(res, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
