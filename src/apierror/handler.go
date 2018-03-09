package apierror

import "auth/src/kind"

type Handler interface {
	Handle(err error) *kind.ApiError
}
