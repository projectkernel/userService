// +build integration
package google

import (
	"testing"
	"encoding/json"
)

// Those Tests can be only manually tested (Thanks Google)

// Use https://developers.google.com/oauthplayground/
// to generate tokens
func TestAccessControl(t *testing.T) {
	t.Skip()
	controller := New("","")
	userData, err := controller.Info("ya29.Glt2BXoQcDmc8ESn0KZ5yxji-CN9e_LNn6V7TR2y0gFFwpYr5P8JdvxIMtVt7wqscWo5chhLrAJFTjSI7nEwKMlFCCVFH3Or7pf6xysah297cMp5aTIECrnI5Qxr")
	if err != nil {
		t.Fatal(err)
	}

	userJson, err := json.Marshal(userData)
	if err != nil {
		t.Fatal("Could not parse json properly")
	}
	t.Log(string(userJson))
}

func TestExchange(t *testing.T) {
	t.Skip()
	controller := New("","")
	access, refresh, err := controller.Auth("4/AACYDX1ww89brEejrM7LmmEpPpG6kq3ttKmGsN6zHo2dyihW2mLrvVFpo4ddCCSxgagYoBLeWY1eWVRl_WlKLIg")
	if err != nil {
		t.Fatal(err)
	}
	t.Log("AcessToken: " + access)
	t.Log("RefreshToken: " + refresh)
}