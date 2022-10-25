package clients

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/tesis/user-ms/config"
	"github.com/tesis/user-ms/errors"
	"io/ioutil"
	"net/http"
)

var ExternalMs = new(config.ExternalMs).Init()

func InterfaceToReader(request interface{}) (*bytes.Buffer, error) {
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(request)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &buf, nil
}

func ValidateHttpResponse(response *http.Response, responseBody interface{}) error {
	if response.StatusCode >= 400 {
		var errorResponse errors.ErrorResponse
		res, err := BodyToInterface(response, &errorResponse)
		if err != nil {
			return err
		}
		if errorResponse.Validate() {
			return &errorResponse
		}
		fmt.Printf("Error in request %s, response, body: %s", response.Request.RequestURI, string(res))
		return &errors.InternalErrorClient
	}
	BodyToInterface(response, &responseBody)
	return nil
}

func BodyToInterface(response *http.Response, bodyInterface interface{}) ([]byte, error) {
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return []byte{}, err
	}
	err = json.Unmarshal(body, bodyInterface)
	if err != nil {
		return []byte{}, err
	}
	return body, nil
}
