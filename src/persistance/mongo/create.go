package mongo

import (
	"auth/src/pojo"
	"log"
	"gopkg.in/mgo.v2/bson"

)

func (db Mongo) Create(data *pojo.User) (user *pojo.User, err error) {
	err = db.collection.Insert(data)
	if err != nil {
		log.Fatal(err)
		return user, err
	}
	user = &pojo.User {}
	err = db.collection.Find(bson.M{}).One(user)
	if err != nil {
		log.Fatal(err)
		return user, err
	}
	return user, nil
}
