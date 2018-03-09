package data

import "auth/src/kind"

type SocialProvider interface {
	Exchange(authCode string) (accessToken string, refreshToken string, err error)
	Info(accessToken string) (user *kind.User, err error)
}
