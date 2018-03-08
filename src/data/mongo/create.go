package mongo

import (
	"auth/src/pojo"

)

func (db Mongo) Create(data *pojo.User) (err error) {
	err = db.collection.Insert(data)
	if err != nil {
		return err
	}
	return nil
}
