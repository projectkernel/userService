package mongo

import (
	"auth/src/pojo"
	"gopkg.in/mgo.v2/bson"
	"log"
)

func (db Mongo) Find(key string) (user *pojo.User, err error) {
	user = &pojo.User {}
	err = db.collection.Find(bson.M{}).One(user)
	if err != nil {
		log.Fatal(err)
		return user, err
	}
	return user, nil
}