package data

import "auth/src/kind"

type Validator interface {
	Validate(user *kind.User) (*kind.User, error)
}

type Cleaner interface {
	Clean(user *kind.User) (*kind.User, error)
}