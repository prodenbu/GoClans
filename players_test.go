package GoClans

import (
	"reflect"
	"testing"
)

func TestGetPlayer(t *testing.T) {
	s := getSession()

	player, err := s.GetPlayer(playerTag)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	if reflect.DeepEqual(player, Player{}) {
		t.Fail()
	}
}

func TestVerifyPlayer(t *testing.T) {
	s := getSession()

	player, err := s.VerifyPlayer(playerTag, "testoken") //should fail
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	if reflect.DeepEqual(player, Player{}) {
		t.Fail()
	}
}
