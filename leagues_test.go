package GoClans

import (
	"reflect"
	"testing"
)

func TestGetLeagues(t *testing.T) {
	s := getSession()

	leagues, err := s.GetLeagues(GetLeaguesRequest{})
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	if reflect.DeepEqual(leagues, []League{}) {
		t.Fail()
	}
}

func TestGetLeagueSeasons(t *testing.T) {
	s := getSession()

	leagues, err := s.GetLeagueSeasons(GetLeagueSeasonsRequest{LeagueId: "29000022"})
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	if reflect.DeepEqual(leagues, []League{}) {
		t.Fail()
	}
}

func TestGetLeaguesSeasonRanking(t *testing.T) {
	s := getSession()

	leagues, err := s.GetLeaguesSeasonRanking(GetLeaguesSeasonRankingRequest{LeagueId: "29000022", SeasonId: "2015-07"})
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	if reflect.DeepEqual(leagues, []PlayerRanking{}) {
		t.Fail()
	}
}

func TestGetLeagueById(t *testing.T) {
	s := getSession()

	leagues, err := s.GetLeagueById("29000022")
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	if reflect.DeepEqual(leagues, League{}) {
		t.Fail()
	}
}

func TestGetWarLeagueById(t *testing.T) {
	s := getSession()

	leagues, err := s.GetWarLeagueById("48000001")
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	if reflect.DeepEqual(leagues, WarLeague{}) {
		t.Fail()
	}
}

func TestGetWarLeagues(t *testing.T) {
	s := getSession()

	leagues, err := s.GetWarLeagues(GetWarLeaguesRequest{})
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	if reflect.DeepEqual(leagues, []WarLeague{}) {
		t.Fail()
	}
}
