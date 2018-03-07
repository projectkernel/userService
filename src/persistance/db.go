package persistance

import "auth/src/pojo"

type DB interface {
	Create(data *pojo.User)(user *pojo.User, err error)
}