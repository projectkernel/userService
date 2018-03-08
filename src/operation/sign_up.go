package operation

import (
	"auth/src/pojo"
	"errors"
)

func (model Persistent) SignUp(userData *pojo.User) (user *pojo.User, err error) {
	err = model.db.Create(userData)
	if err != nil {
		return user, errors.New("could not save user data")
	}
	user, err = model.db.Find("TODO")
	return user,nil
}