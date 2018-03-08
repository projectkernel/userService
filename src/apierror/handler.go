package apierror

import "auth/src/pojo"

type Handler interface {
	Handle(err error) *pojo.ApiError
}
