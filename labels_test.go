package GoClans

import (
	"reflect"
	"testing"
)

func TestGetClanLabels(t *testing.T) {
	s := getSession()

	labels, err := s.GetClanLabels()
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	if reflect.DeepEqual(labels, []Label{}) {
		t.Fail()
	}
}

func TestGetPlayerLabels(t *testing.T) {
	s := getSession()

	labels, err := s.GetPlayerLabels()
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	if reflect.DeepEqual(labels, []Label{}) {
		t.Fail()
	}
}
