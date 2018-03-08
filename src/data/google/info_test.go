// +build integration
package google

import (
	"encoding/json"
	"testing"
)

// Use https://developers.google.com/oauthplayground/
// to generate tokens
func TestAccessControl(t *testing.T) {
	t.Skip()
	userData, err := Info("ya29.Glt2BXoQcDmc8ESn0KZ5yxji-CN9e_LNn6V7TR2y0gFFwpYr5P8JdvxIMtVt7wqscWo5chhLrAJFTjSI7nEwKMlFCCVFH3Or7pf6xysah297cMp5aTIECrnI5Qxr")
	if err != nil {
		t.Fatal(err)
	}

	userJson, err := json.Marshal(userData)
	if err != nil {
		t.Fatal("Could not parse json properly")
	}
	t.Log(string(userJson))
}