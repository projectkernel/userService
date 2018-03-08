package apierror

import "auth/src/pojo"

type SimpleHandler struct {}

// TODO Implement
func (SimpleHandler) Handle(err error) *pojo.ApiError {
	return &pojo.ApiError{
		Message: "undefined error occurred",
		Code: UNKNOWN,
	}
}
