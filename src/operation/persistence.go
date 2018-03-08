package operation

import "auth/src/data"

type Persistent struct {
	db data.DB
}

func NewPersistent(db data.DB) *Persistent {
	return &Persistent{
		db: db,
	}
}