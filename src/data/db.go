package data

import "auth/src/kind"

type DB interface {
	Create(data *kind.User)(user *kind.User, err error)
	Find(key string) (user *kind.User, err error)
	Upsert(data *kind.User)(user *kind.User, err error)
}