package mongo

import (
	"gopkg.in/mgo.v2"
	"auth/src/pojo"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type Mongo struct {
	collection *mgo.Collection
}

func NewMongo(session *mgo.Session) (*Mongo) {
	return &Mongo{
		collection: session.DB("test").C("users"),
	}
}

func (db Mongo) Create(data *pojo.User) (err error) {
	err = db.collection.Insert(data)
	if err != nil {
		return err
	}
	return nil
}

func (db Mongo) Find(key string) (user *pojo.User, err error) {
	user = &pojo.User {}
	err = db.collection.Find(bson.M{}).One(user)
	if err != nil {
		log.Fatal(err)
		return user, err
	}
	return user, nil
}