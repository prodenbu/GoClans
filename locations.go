package GoClans

import (
	"encoding/json"
	"fmt"
	"net/url"
)

/*
Get clan rankings for a specific location

parameters:

required locationId string: Identifier of the location to retrieve.

optional limit int: Limit the number of items returned in the response.

optional after string: Return only items that occur after this marker. Before marker can be found from the response,
inside the 'paging' property. Note that only after or before can be specified for a request, not both.

optional before string: Return only items that occur before this marker. Before marker can be found from the response,
inside the 'paging' property. Note that only after or before can be specified for a request, not both.
*/
func (s *Session) GetLocationClanRankings(r GetLocationRankingsRequest) ([]ClanRanking, error) {
	ranking := ClanRankingList{}

	url, err := url.Parse(fmt.Sprintf("locations/%s/rankings/clans", r.LocationId))
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
Get player rankings for a specific location

parameters:

required locationId string: Identifier of the location to retrieve.

optional limit int: Limit the number of items returned in the response.

optional after string: Return only items that occur after this marker. Before marker can be found from the response,
inside the 'paging' property. Note that only after or before can be specified for a request, not both.

optional before string: Return only items that occur before this marker. Before marker can be found from the response,
inside the 'paging' property. Note that only after or before can be specified for a request, not both.
*/
func (s *Session) GetLocationPlayerRankings(r GetLocationRankingsRequest) ([]PlayerRanking, error) {
	ranking := PlayerRankingList{}

	url, err := url.Parse(fmt.Sprintf("locations/%s/rankings/players", r.LocationId))
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
Get player versus rankings for a specific location

parameters:

required locationId string: Identifier of the location to retrieve.

optional limit int: Limit the number of items returned in the response.

optional after string: Return only items that occur after this marker. Before marker can be found from the response,
inside the 'paging' property. Note that only after or before can be specified for a request, not both.

optional before string: Return only items that occur before this marker. Before marker can be found from the response,
inside the 'paging' property. Note that only after or before can be specified for a request, not both.
*/
func (s *Session) GetLocationPlayerVersusRankings(r GetLocationRankingsRequest) ([]PlayerVersusRanking, error) {
	ranking := PlayerVersusRankingList{}

	url, err := url.Parse(fmt.Sprintf("locations/%s/rankings/players-versus", r.LocationId))
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
Get clan versus rankings for a specific location

parameters:

required locationId string: Identifier of the location to retrieve.

optional limit int: Limit the number of items returned in the response.

optional after string: Return only items that occur after this marker. Before marker can be found from the response,
inside the 'paging' property. Note that only after or before can be specified for a request, not both.

optional before string: Return only items that occur before this marker. Before marker can be found from the response,
inside the 'paging' property. Note that only after or before can be specified for a request, not both.
*/
func (s *Session) GetLocationClanVersusRankings(r GetLocationRankingsRequest) ([]ClanVersusRanking, error) {
	ranking := ClanVersusRankingList{}

	url, err := url.Parse(fmt.Sprintf("locations/%s/rankings/clans-versus", r.LocationId))
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
List locations

parameters:

optional limit int: Limit the number of items returned in the response.

optional after string: Return only items that occur after this marker. Before marker can be found from the response,
inside the 'paging' property. Note that only after or before can be specified for a request, not both.

optional before string: Return only items that occur before this marker. Before marker can be found from the response,
inside the 'paging' property. Note that only after or before can be specified for a request, not both.
*/
func (s *Session) GetLocations(r GetLocationsRequest) ([]PlayerVersusRanking, error) {
	ranking := PlayerVersusRankingList{}

	url, err := url.Parse("locations")
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
Get information about specific location

parameters:

required locationId string: Identifier of the location to retrieve.

optional limit int: Limit the number of items returned in the response.

optional after string: Return only items that occur after this marker. Before marker can be found from the response,
inside the 'paging' property. Note that only after or before can be specified for a request, not both.

optional before string: Return only items that occur before this marker. Before marker can be found from the response,
inside the 'paging' property. Note that only after or before can be specified for a request, not both.
*/
func (s *Session) GetLocationById(locationId string) (Location, error) {
	ranking := Location{}

	body, err := s.get(fmt.Sprintf("locations/%s", locationId))
	if err != nil {
		return ranking, err
	}
	err = json.Unmarshal(body, &ranking)
	if err != nil {
		return ranking, err
	}
	return ranking, nil
}
