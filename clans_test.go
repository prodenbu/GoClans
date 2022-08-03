package GoClans

import (
	"reflect"
	"testing"
)

const (
	key            string = ""
	clanTag        string = ""
	leagueWarTag   string = ""
	playerTag      string = ""
	clanSearchName string = ""
)

func getSession() *Session {
	return &Session{ApiKey: key}

}

func TestGetLeagueGroup(t *testing.T) {
	s := getSession()

	LeagueGroup, err := s.GetLeagueGroup(clanTag)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	if reflect.DeepEqual(LeagueGroup, ClanWarLeagueGroup{}) {
		t.Fail()
	}
}

func TestGetLeagueWar(t *testing.T) {
	s := getSession()

	LeagueGroup, err := s.GetLeagueWar(leagueWarTag)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	if reflect.DeepEqual(LeagueGroup, ClanWarLeagueGroup{}) {
		t.Fail()
	}
}

func TestGetWarLog(t *testing.T) {
	s := getSession()

	log, err := s.GetWarLog(ClanWarLogRequest{ClanTag: clanTag})
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	if reflect.DeepEqual(log, ClanWarLog{}) {
		t.Fail()
	}
}

func TestSearchClans(t *testing.T) {
	s := getSession()

	log, err := s.SearchClans(SearchClansRequest{Name: clanSearchName})
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	if reflect.DeepEqual(log, ClanWarLog{}) {
		t.Fail()
	}
}

func TestGetCurrentWar(t *testing.T) {
	s := getSession()

	war, err := s.GetCurrentWar(clanTag)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	if reflect.DeepEqual(war, ClanWar{}) {
		t.Fail()
	}
}

func TestGetClan(t *testing.T) {
	s := getSession()

	clan, err := s.GetClan(clanTag)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	if reflect.DeepEqual(clan, Clan{}) {
		t.Fail()
	}
}

func TestGetClanMember(t *testing.T) {
	s := getSession()

	member, err := s.GetClanMember(ClanMemberRequest{ClanTag: clanTag})
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	if reflect.DeepEqual(member, []ClanMember{}) {
		t.Fail()
	}
}
