package mongo

import "gopkg.in/mgo.v2"

type Mongo struct {
	collection *mgo.Collection
}

func NewMongo(session *mgo.Session) (*Mongo) {
	return &Mongo{
		collection: session.DB("test").C("users"),
	}
}
