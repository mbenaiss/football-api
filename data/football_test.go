package data

import "testing"

func TestGetDataFromApi(t *testing.T) {
	matches, err := GetDataFromApi()
	if err != nil {
		t.Error(err)
	}
	if len(matches) == 0 {
		t.Errorf("expected > 0 got 0")
	}
}
