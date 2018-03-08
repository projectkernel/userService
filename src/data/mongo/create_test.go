package mongo

import (
	"testing"
	"auth/src/pojo"
	"fmt"
	"log"
	"encoding/json"
)

var crud *Mongo

func TestMain(m *testing.M) {
	session, _ := Get([]string{"localhost"})
	crud = NewMongo(session)
	m.Run()
}

func TestCreate(t *testing.T) {
	user := &pojo.User{
		Name: "Test",
		Email: "test@email.com",
		ImageURL: "image.com",
		Language: "pt-BR",
		Provider: "google",
		RefreshToken: "randomtoken",
	}
	result, err := crud.Create(user)
	if err != nil {
		log.Fatal(err)
	}
	x, err := json.Marshal(result)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Print(string(x))
}
