package data

import (
	"auth/src/kind"
	"auth/src/util"
)

type Persistence struct {
	validator Validator
	db     DB
	cleaner Cleaner
}

func NewPersistence(db DB) *Persistence {
	return &Persistence{
		db:     db,
	}
}

type Result struct {
	User *kind.User `json:"user"`
	Token string     `json:"token"`
}

func (pst Persistence) Upsert(user *kind.User) (*kind.User, error) {
	return pst.do(pst.db.Upsert, user)
}

func (pst Persistence) do(op kind.UserOp, user *kind.User) (*kind.User, error) {
	return util.PipeUserOps([]kind.UserOp {
		pst.validator.Validate,
		op,
		pst.cleaner.Clean,
	}, user)
}