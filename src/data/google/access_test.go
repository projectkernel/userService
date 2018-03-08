// +build integration
package google

import (
	"testing"
)

func TestExchange(t *testing.T) {
	t.Skip()
	controller := NewProvider("","")
	access, refresh, err := controller.Exchange("4/AACYDX1ww89brEejrM7LmmEpPpG6kq3ttKmGsN6zHo2dyihW2mLrvVFpo4ddCCSxgagYoBLeWY1eWVRl_WlKLIg")
	if err != nil {
		t.Fatal(err)
	}
	t.Log("AcessToken: " + access)
	t.Log("RefreshToken: " + refresh)
}