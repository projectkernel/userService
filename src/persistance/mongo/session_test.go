// +build integration
package mongo

import (
	"testing"
)

func TestGetLocally(t *testing.T) {
	session , err := Get([]string{"localhost"})
	if err != nil {
		t.Fatal(err)
	}
	defer session.Close()
}

// Integration test
func TestGetMongoAtlas(t *testing.T) {
	session , err := GetWithCredentials(
		[]string{
			"cluster0-shard-00-00-lf760.mongodb.net",
			"cluster0-shard-00-01-lf760.mongodb.net",
			"cluster0-shard-00-02-lf760.mongodb.net",
		},
		"admin",
		"admin",
		"",
		true)
	if err != nil {
		t.Fatal(err)
	}
	defer session.Close()
}