package signin

import (
	"testing"
	"auth/src/persistance"
	"auth/src/persistance/mongo"
)

var db persistance.DB := mongo.NewMongo()
var model Model := NewModel(nil)

func TestMain(m *testing.M) {

}