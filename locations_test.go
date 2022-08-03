package GoClans

import (
	"log"
	"reflect"
	"testing"
)

func TestGetLocationClanRankings(t *testing.T) {
	s := getSession()

	ranking, err := s.GetLocationClanRankings(GetLocationRankingsRequest{LocationId: "32000007"})

	if err != nil {
		t.Log(err)
		t.Fail()
	}
	if reflect.DeepEqual(ranking, []ClanRanking{}) {
		t.Logf("is empty: %+v", ranking)
		t.Fail()
	}
}

func TestGetLocationPlayerRankings(t *testing.T) {
	s := getSession()

	ranking, err := s.GetLocationPlayerRankings(GetLocationRankingsRequest{LocationId: "32000007"})

	if err != nil {
		t.Log(err)
		t.Fail()
	}
	if reflect.DeepEqual(ranking, []PlayerRanking{}) {
		t.Logf("is empty: %+v", ranking)
		t.Fail()
	}
}

func TestGetLocationPlayerVersusRankings(t *testing.T) {
	s := getSession()

	ranking, err := s.GetLocationPlayerVersusRankings(GetLocationRankingsRequest{LocationId: "32000007"})

	if err != nil {
		t.Log(err)
		t.Fail()
	}
	if reflect.DeepEqual(ranking, []PlayerVersusRanking{}) {
		t.Logf("is empty: %+v", ranking)
		t.Fail()
	}
}

func TestGetLocationClansVersusRankings(t *testing.T) {
	s := getSession()

	ranking, err := s.GetLocationClanVersusRankings(GetLocationRankingsRequest{LocationId: "32000007"})

	if err != nil {
		t.Log(err)
		t.Fail()
	}
	log.Printf("%+v", ranking[0])
	if reflect.DeepEqual(ranking, []ClanVersusRanking{}) {
		t.Logf("is empty: %+v", ranking)
		t.Fail()
	}
}
func TestGetLocations(t *testing.T) {
	s := getSession()

	ranking, err := s.GetLocations(GetLocationsRequest{})

	if err != nil {
		t.Log(err)
		t.Fail()
	}
	if reflect.DeepEqual(ranking, []PlayerVersusRanking{}) {
		t.Logf("is empty: %+v", ranking)
		t.Fail()
	}
}

func TestGetLocationById(t *testing.T) {
	s := getSession()

	ranking, err := s.GetLocationById("32000007")

	if err != nil {
		t.Log(err)
		t.Fail()
	}
	if reflect.DeepEqual(ranking, []PlayerVersusRanking{}) {
		t.Logf("is empty: %+v", ranking)
		t.Fail()
	}
}
