package apierror

import "auth/src/kind"

type SimpleHandler struct {}

// TODO Implement
func (SimpleHandler) Handle(err error) *kind.ApiError {
	return &kind.ApiError{
		Message: "undefined error occurred",
		Code: UNKNOWN,
	}
}
