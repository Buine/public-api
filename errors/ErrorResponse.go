package errors

import (
	"fmt"
	"strings"
)

type ErrorResponse struct {
	Code     string   `json:"code"`
	Messages []string `json:"messages"`
	Status   int      `json:"-"`
}

func (current *ErrorResponse) Error() string {
	return fmt.Sprintf("[code=%s] messages=%s, status=%d", current.Code, strings.Join(current.Messages, ", "), current.Status)
}

func (current *ErrorResponse) Validate() bool {
	return !(current.Code == "" && len(current.Code) == 0)
}

var InternalErrorClient = ErrorResponse{
	Code: "INTERNAL_ERROR",
	Messages: []string{
		"Client failed",
	},
	Status: 500,
}

var SessionExpired = ErrorResponse{
	Code: "SESSION_EXPIRED",
	Messages: []string{
		"Session expired",
	},
	Status: 400,
}

var BadRequest = ErrorResponse{
	Code: "BAD_REQUEST",
	Messages: []string{
		"Bad request",
	},
	Status: 400,
}

var Unauthenticated = ErrorResponse{
	Code: "UNAUTHENTICATED",
	Messages: []string{
		"Unauthenticated",
	},
	Status: 400,
}
