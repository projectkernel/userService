package controller

import (
	"auth/src/pipeline"
	"auth/src/data"
	"auth/src/pojo"
)

type ProviderPersistence struct {
	db     data.DB
	prov   *pipeline.Provider
}

func NewProviderPersistence(prov *pipeline.Provider, db data.DB) *ProviderPersistence {
	return &ProviderPersistence{
		prov:   prov,
		db:     db,
	}
}

type Result struct {
	User *pojo.User `json:"user"`
	Token string `json:"token"`
}

// Return type based on Formatter Template
func (ctrl ProviderPersistence) SocialSignUp(token string) (res *Result, err error) {
	user, err := ctrl.prov.GetUserFromProvider(token)
	if err != nil {
		return res, err
	}
	err = ctrl.db.Create(user)
	if err != nil {
		return res, err
	}
	// Erases refreshToken after it is already saved
	user.RefreshToken = ""
	user, err =ctrl.db.Find("TODO")
	if err != nil {
		return res, err
	}
	// TODO Use JWT
	res = &Result{
		User:user,
		Token:"",
	}
	return res, nil
}