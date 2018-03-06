package google

import (
	"auth/src/pojo"
	"fmt"
	"auth/src/create"
	"errors"
)

func SignIn(authCode string) (user pojo.User, accessToken string, err error) {
	accessToken, refreshToken, err := auth(authCode)
	fmt.Print(refreshToken)
	if err != nil {
		return user, "", errors.New("authentication failed")
	}
	userData, err := info(accessToken)
	user = create.Create(userData)
	return user, accessToken, nil
}