package data

import "auth/src/pojo"

type DB interface {
	Create(data *pojo.User)(err error)
	Find(key string) (user *pojo.User, err error)
}