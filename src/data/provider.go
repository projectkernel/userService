package data

import "auth/src/pojo"

type Provider interface {
	Exchange(authCode string) (accessToken string, refreshToken string, err error)
	Info(accessToken string) (user *pojo.User, err error)
}
