package GoClans

import (
	"encoding/json"
	"fmt"
	"net/url"
)

/*
List leagues

parameters:

optional limit int: Limit the number of items returned in the response.

optional after string: Return only items that occur after this marker. Before marker can be found from the response,
inside the 'paging' property. Note that only after or before can be specified for a request, not both.

optional before string: Return only items that occur before this marker. Before marker can be found from the response,
inside the 'paging' property. Note that only after or before can be specified for a request, not both.
*/
func (s *Session) GetLeagues(r GetLeaguesRequest) ([]League, error) {
	leagues := LeagueList{}

	url, err := url.Parse("leagues")
	if err != nil {
		return leagues.Items, err
	}
	params := url.Query()

	if r.After != "" {
		params.Add("after", r.After)
	}
	if r.Before != "" {
		params.Add("before", r.Before)
	}
	if r.Limit > 0 {
		params.Add("limit", fmt.Sprintf("%d", r.Limit))
	}

	url.RawQuery = params.Encode()

	body, err := s.get(url.String())
	if err != nil {
		return leagues.Items, err
	}
	err = json.Unmarshal(body, &leagues)
	if err != nil {
		return leagues.Items, err
	}
	return leagues.Items, nil
}

/*
Get league seasons. Note that league season information is available only for Legend League.

parameters:

required leagueId string: Identifier of the league

optional limit int: Limit the number of items returned in the response.

optional after string: Return only items that occur after this marker. Before marker can be found from the response,
inside the 'paging' property. Note that only after or before can be specified for a request, not both.

optional before string: Return only items that occur before this marker. Before marker can be found from the response,
inside the 'paging' property. Note that only after or before can be specified for a request, not both.
*/
func (s *Session) GetLeagueSeasons(r GetLeagueSeasonsRequest) ([]LeagueSeason, error) {
	seasons := LeagueSeasonList{}

	url, err := url.Parse(fmt.Sprintf("leagues/%s/seasons", r.LeagueId))
	if err != nil {
		return seasons.Items, err
	}
	params := url.Query()

	if r.After != "" {
		params.Add("after", r.After)
	}
	if r.Before != "" {
		params.Add("before", r.Before)
	}
	if r.Limit > 0 {
		params.Add("limit", fmt.Sprintf("%d", r.Limit))
	}

	url.RawQuery = params.Encode()

	body, err := s.get(url.String())
	if err != nil {
		return seasons.Items, err
	}
	err = json.Unmarshal(body, &seasons)
	if err != nil {
		return seasons.Items, err
	}
	return seasons.Items, nil
}

/*
Get league season rankings. Note that league season information is available only for Legend League.

Parameters:

required leagueId string: Identifier of the league.

required seasonId string: Identifier of the season.

optional limit int: Limit the number of items returned in the response.

optional after string: Return only items that occur after this marker. Before marker can be found from the response,
inside the 'paging' property. Note that only after or before can be specified for a request, not both.

optional before string: Return only items that occur before this marker. Before marker can be found from the response,
inside the 'paging' property. Note that only after or before can be specified for a request, not both.
*/
func (s *Session) GetLeaguesSeasonRanking(r GetLeaguesSeasonRankingRequest) ([]PlayerRanking, error) {
	ranking := PlayerRankingList{}

	url, err := url.Parse(fmt.Sprintf("leagues/%s/seasons/%s", r.LeagueId, r.SeasonId))
	if err != nil {
		return ranking.Items, err
	}
	params := url.Query()

	if r.After != "" {
		params.Add("after", r.After)
	}
	if r.Before != "" {
		params.Add("before", r.Before)
	}
	if r.Limit > 0 {
		params.Add("limit", fmt.Sprintf("%d", r.Limit))
	}

	url.RawQuery = params.Encode()

	body, err := s.get(url.String())
	if err != nil {
		return ranking.Items, err
	}
	err = json.Unmarshal(body, &ranking)
	if err != nil {
		return ranking.Items, err
	}
	return ranking.Items, nil
}

/*
Get league information.

parameters:

required leagueId string: Identifier of the league
*/
func (s *Session) GetLeagueById(leagueId string) (League, error) {
	league := League{}
	body, err := s.get(fmt.Sprintf("leagues/%s", leagueId))
	if err != nil {
		return league, err
	}
	err = json.Unmarshal(body, &league)
	if err != nil {
		return league, err
	}
	return league, nil
}

/*
Get war league information.

parameters:

required leagueId string: Identifier of the league
*/
func (s *Session) GetWarLeagueById(leagueId string) (WarLeague, error) {
	league := WarLeague{}
	body, err := s.get(fmt.Sprintf("warleagues/%s", leagueId))
	if err != nil {
		return league, err
	}
	err = json.Unmarshal(body, &league)
	if err != nil {
		return league, err
	}
	return league, nil
}

/*
List war leagues.

parameters:

optional limit int: Limit the number of items returned in the response.

optional after string: Return only items that occur after this marker. Before marker can be found from the response,
inside the 'paging' property. Note that only after or before can be specified for a request, not both.

optional before string: Return only items that occur before this marker. Before marker can be found from the response,
inside the 'paging' property. Note that only after or before can be specified for a request, not both.
*/
func (s *Session) GetWarLeagues(r GetWarLeaguesRequest) ([]WarLeague, error) {
	leagues := WarLeagueList{}

	url, err := url.Parse("warleagues")
	if err != nil {
		return leagues.Items, err
	}
	params := url.Query()

	if r.After != "" {
		params.Add("after", r.After)
	}
	if r.Before != "" {
		params.Add("before", r.Before)
	}
	if r.Limit > 0 {
		params.Add("limit", fmt.Sprintf("%d", r.Limit))
	}

	url.RawQuery = params.Encode()

	body, err := s.get(url.String())
	if err != nil {
		return leagues.Items, err
	}
	err = json.Unmarshal(body, &leagues)
	if err != nil {
		return leagues.Items, err
	}
	return leagues.Items, nil
}
