package clients

import (
	"fmt"
	"github.com/gojek/heimdall/v7/httpclient"
	"github.com/tesis/user-ms/constants/routes"
	"github.com/tesis/user-ms/requests"
	"github.com/tesis/user-ms/responses"
)

func CreateUser(userRequest requests.UserRequest) (*responses.UserResponse, error) {
	var response responses.UserResponse
	body, err := InterfaceToReader(userRequest)
	if err != nil {
		return nil, err
	}
	client := httpclient.NewClient()
	res, err := client.Post(fmt.Sprintf(routes.CreateUserRoute, ExternalMs.UserMs), body, nil)
	if err != nil {
		return nil, err
	}
	err = ValidateHttpResponse(res, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func GetUser(code string) (*responses.UserResponse, error) {
	var response responses.UserResponse
	client := httpclient.NewClient()
	res, err := client.Get(fmt.Sprintf(routes.GetUser, ExternalMs.UserMs, code), nil)
	if err != nil {
		return nil, err
	}
	err = ValidateHttpResponse(res, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func GetUserLogin(loginRequest requests.UserLoginRequest) (*responses.UserResponse, error) {
	var response responses.UserResponse
	body, err := InterfaceToReader(loginRequest)
	if err != nil {
		return nil, err
	}
	client := httpclient.NewClient()
	res, err := client.Post(fmt.Sprintf(routes.LoginUserRoute, ExternalMs.UserMs), body, nil)
	if err != nil {
		return nil, err
	}
	err = ValidateHttpResponse(res, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
