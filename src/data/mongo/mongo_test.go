package mongo

import (
	"testing"
	"auth/src/kind"
	"github.com/google/go-cmp/cmp"
)

var crud *Mongo
var user = &kind.User{
	Name: "Test",
	Email: "test@email.com",
	ImageURL: "image.com",
	Language: "pt-BR",
	Provider: "google",
	RefreshToken: "randomtoken",
}
var newToken = "randomtoken2"

func TestMain(m *testing.M) {
	session, _ := Get([]string{"localhost"})
	crud = NewMongo(session)
	crud.Drop()
	m.Run()
}

func TestCreate(t *testing.T) {
	_, err := crud.Create(user)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUpsert(t *testing.T) {
	_, err := crud.Create(user)
	if err != nil {
		t.Fatal(err)
	}
	otherUser := user
	otherUser.RefreshToken = newToken
	_, err = crud.Upsert(otherUser)
	if err != nil {
		t.Fatal(err)
	}
	dbUser, err := crud.Find(otherUser.Email)
	if err != nil {
		t.Fatal(err)
	}
	if dbUser.RefreshToken != newToken {
		t.Fatal("Value should have been updated at Upsert")
	}
}

func TestFind(t *testing.T) {
	_, err := crud.Create(user)
	if err != nil {
		t.Fatal(err)
	}
	dbUser, err := crud.Find(user.Email)
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(dbUser, user) {
		t.Fatal("Should be equal")
	}
}

func TestDrop(t *testing.T) {
	err := crud.Drop()
	if err != nil {
		t.Fatal(err)
	}
}