package signin

import (
	"auth/src/pojo"
	"fmt"
	"auth/src/persistance"
	"errors"
)

type Model struct {
	db *persistance.DB
}

func NewModel(db *persistance.DB) *Model {
	return &Model{
		db: db,
	}
}

func (m Model) SignIn(method Method, authCode string) (user *pojo.User, accessToken string, err error) {
	accessToken, refreshToken, err := method.Auth(authCode)
	fmt.Print(refreshToken)
	if err != nil {
		return user, "", errors.New("authentication failed")
	}
	userData, err := method.Info(accessToken)
	if err != nil {
		return user, "", errors.New("could not get required data")
	}
	userData.RefreshToken = refreshToken
	user, err = m.db.Create(userData)
	// Erases refreshToken after it is already saved
	userData.RefreshToken = ""
	if err != nil {
		return user, "", errors.New("could not save user data")
	}
	return user, accessToken, nil
}