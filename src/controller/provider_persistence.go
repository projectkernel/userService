package controller

import (
	"auth/src/pipeline"
	"auth/src/data"
	"auth/src/kind"
	"auth/src/auth"
	"auth/src/util"
)

type ProviderPersistence struct {
	prov   *pipeline.Provider
	validator data.Validator
	db     data.DB
	cleaner data.Cleaner
	encryption auth.Encryption
}

func NewProviderPersistence(prov *pipeline.Provider, db data.DB, encryption auth.Encryption) *ProviderPersistence {
	return &ProviderPersistence{
		prov:   prov,
		db:     db,
		encryption: encryption,
	}
}

type Result struct {
	User *kind.User `json:"user"`
	Token string     `json:"token"`
}

func (ctrl ProviderPersistence) SocialSignUp(token string) (res *Result, err error) {
	user, err := ctrl.prov.GetUserFromProvider(token)
	if err != nil {
		return res, err
	}
	user, err = util.PipeUserOps([]kind.UserOp{
		ctrl.validator.Validate,
		ctrl.db.Upsert,
		ctrl.cleaner.Clean,
	}, user)
	//user, err = ctrl.validator.Validate(user)
	//if err != nil {
	//	return res, err
	//}
	//user, err = ctrl.db.Upsert(user)
	//if err != nil {
	//	return res, err
	//}
	//user = ctrl.cleaner.Clean(user)
	res = &Result{
		User: user,
		Token: ctrl.encryption.Encrypt(user.Email),
	}
	return res, nil
}