package mongo

import (
	"gopkg.in/mgo.v2"
	"auth/src/kind"
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

func (db Mongo) Create(data *kind.User) (*kind.User, error) {
	err := db.collection.Insert(data)
	if err != nil {
		return data, err
	}
	return data, nil
}

func (db Mongo) Find(email string) (user *kind.User, err error) {
	user = &kind.User {}
	err = db.collection.Find(bson.M{"_id": email}).One(user)
	if err != nil {
		log.Fatal(err)
		return user, err
	}
	return user, nil
}

func (db Mongo) Upsert(data *kind.User)(*kind.User, error) {
	info, err := db.collection.Upsert(bson.M{"_id": data.Email}, data)
	if err != nil {
		return data, err
	}
	if info.Matched == 0 {
		// Login for the first time
	}
	return data, nil
}

func (db Mongo) Drop() (error) {
	return db.collection.DropCollection()
}
