package services

import (
	"github.com/tesis/user-ms/clients"
	"github.com/tesis/user-ms/requests"
	"github.com/tesis/user-ms/responses"
)

func CreateUser(userRequest requests.UserRequest) (*responses.UserResponse, error) {
	res, err := clients.CreateUser(userRequest)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func GetUser(code string) (*responses.UserResponse, error) {
	res, err := clients.GetUser(code)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func GetUserLogin(request requests.UserLoginRequest) (*responses.UserResponse, error) {
	res, err := clients.GetUserLogin(request)
	if err != nil {
		return nil, err
	}
	return res, nil
}
