package GoClans

import (
	"reflect"
	"testing"
)

func TestGoldPass(t *testing.T) {
	s := getSession()

	pass, err := s.GetCurrentGoldpassSeason()
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	if reflect.DeepEqual(pass, GoldPassSeason{}) {
		t.Fail()
	}
}
