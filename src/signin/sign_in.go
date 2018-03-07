package signin

import (
	"auth/src/pojo"
)

type Method interface {
	Auth(authCode string) (accessToken string, refreshToken string, err error)
	Info(accessToken string) (user *pojo.User, err error)
}