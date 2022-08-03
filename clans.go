package GoClans

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

/*
Retrieve information about clan's current clan war league group

Parameters:

required clanTag string: Tag of the clan.
*/
func (s *Session) GetLeagueGroup(clantag string) (ClanWarLeagueGroup, error) {
	leageuGroup := ClanWarLeagueGroup{}
	body, err := s.get(fmt.Sprintf("clans/%s/currentwar/leaguegroup", strings.Replace(clantag, "#", "%23", 1)))
	if err != nil {
		return leageuGroup, err
	}
	err = json.Unmarshal(body, &leageuGroup)
	if err != nil {
		return leageuGroup, err
	}
	return leageuGroup, nil
}

/*
Retrieve information about individual clan war league war

Parameters:

required warTag string: Tag of the war.
*/
func (s *Session) GetLeagueWar(wartag string) (ClanWarLeagueGroup, error) {
	leageuGroup := ClanWarLeagueGroup{}
	body, err := s.get(fmt.Sprintf("clanwarleagues/wars/%s", strings.Replace(wartag, "#", "%23", 1)))
	if err != nil {
		return leageuGroup, err
	}
	err = json.Unmarshal(body, &leageuGroup)
	if err != nil {
		return leageuGroup, err
	}
	return leageuGroup, nil
}

/*
Retrieve clan's clan war log

Parameters:

required clanTag string: Tag of the clan.

optional limit int: Limit the number of items returned in the response.

optional after string: Return only items that occur after this marker. Before marker can be found from the response, inside the 'paging' property. Note that only after or before can be specified for a request, not both.

optional before string: Return only items that occur before this marker. Before marker can be found from the response, inside the 'paging' property. Note that only after or before can be specified for a request, not both.
*/
func (s *Session) GetWarLog(r ClanWarLogRequest) (ClanWarLog, error) {
	warLog := ClanWarLog{}
	url, err := url.Parse(fmt.Sprintf("clans/%s/warlog", strings.Replace(r.ClanTag, "#", "%23", 1)))
	if err != nil {
		return warLog, err
	}
	params := url.Query()

	if r.Limit > 0 {
		params.Add("limit", fmt.Sprintf("%d", r.Limit))
	}
	if r.After > 0 {
		params.Add("after", fmt.Sprintf("%d", r.After))
	}
	if r.Before > 0 {
		params.Add("before", fmt.Sprintf("%d", r.Before))
	}
	url.RawQuery = params.Encode()

	body, err := s.get(url.String())
	if err != nil {
		return warLog, err
	}
	err = json.Unmarshal(body, &warLog)
	if err != nil {
		return warLog, err
	}
	return warLog, nil
}

/*
Search all clans by name and/or filtering the results using various criteria. At least one filtering criteria must be
defined and if name is used as part of search, it is required to be at least three characters long. It is not possible
to specify ordering for results so clients should not rely on any specific ordering as that may change in the future
releases of the API.

parameters:

optional name string: Search clans by name. If name is used as part of search query, it needs to be at least three characters long.
Name search parameter is interpreted as wild card search, so it may appear anywhere in the clan name.

optional warFrequency string: Filter by clan war frequency

optional locationId int: Filter by clan location identifier. For list of available locations, refer to getLocations operation.

optional minMembers int: Filter by minimum number of clan members.

optional maxMembers int: Filter by maximum number of clan members.

optional minClanPoints int: Filter by minimum amount of clan points.

optional minClanLevel int: Filter by minimum clan level.

optional limit int: Limit the number of items returned in the response.

optional after string: Return only items that occur after this marker. Before marker can be found from the response,
inside the 'paging' property. Note that only after or before can be specified for a request, not both.

optional before string: Return only items that occur before this marker. Before marker can be found from the response,
inside the 'paging' property. Note that only after or before can be specified for a request, not both.

optional labelIds string: Comma separatered list of label IDs to use for filtering results.
*/
func (s *Session) SearchClans(r SearchClansRequest) (ClanList, error) {
	clans := ClanList{}
	url, err := url.Parse("clans")
	if err != nil {
		return clans, err
	}
	params := url.Query()

	if r.Name != "" {
		params.Add("name", r.Name)
	}
	if r.WarFrequency != "" {
		params.Add("warFrequency", r.WarFrequency)
	}
	if r.LocationId > 0 {
		params.Add("locationId", fmt.Sprintf("%d", r.LocationId))
	}
	if r.MinMembers > 0 {
		params.Add("minMembers", fmt.Sprintf("%d", r.MinMembers))
	}
	if r.MaxMembers > 0 {
		params.Add("maxMembers", fmt.Sprintf("%d", r.MaxMembers))
	}
	if r.MinClanPoints > 0 {
		params.Add("minClanPoints", fmt.Sprintf("%d", r.MinClanPoints))
	}
	if r.MinClanLevel > 0 {
		params.Add("minClanLevel", fmt.Sprintf("%d", r.MinClanLevel))
	}
	if r.Limit > 0 {
		params.Add("limit", fmt.Sprintf("%d", r.Limit))
	}
	if r.After != "" {
		params.Add("after", r.After)
	}
	if r.Before != "" {
		params.Add("before", r.Before)
	}
	if r.LabelIds != "" {
		params.Add("labelIds", r.LabelIds)
	}

	url.RawQuery = params.Encode()

	body, err := s.get(url.String())
	if err != nil {
		return clans, err
	}
	err = json.Unmarshal(body, &clans)
	if err != nil {
		return clans, err
	}
	return clans, nil
}

/*
Retrieve information about clan's current clan war

parameters:

required clanTag string: Tag of the clan.
*/
func (s *Session) GetCurrentWar(clanTag string) (ClanWar, error) {
	war := ClanWar{}
	body, err := s.get(fmt.Sprintf("clans/%s/currentwar", strings.Replace(clanTag, "#", "%23", 1)))
	if err != nil {
		return war, err
	}
	err = json.Unmarshal(body, &war)
	if err != nil {
		return war, err
	}
	return war, nil
}

/*
Get information about a single clan by clan tag. Clan tags can be found using clan search operation.

parameters:

required clanTag string: Tag of the clan.
*/
func (s *Session) GetClan(clanTag string) (Clan, error) {
	clan := Clan{}
	body, err := s.get(fmt.Sprintf("clans/%s", strings.Replace(clanTag, "#", "%23", 1)))
	if err != nil {
		return clan, err
	}
	err = json.Unmarshal(body, &clan)
	if err != nil {
		return clan, err
	}
	return clan, nil
}

/*
List clan members.

parameters:

required clanTag string: Tag of the clan.

optional limit int: Limit the number of items returned in the response.

optional after string: Return only items that occur after this marker. Before marker can be found from the response,
inside the 'paging' property. Note that only after or before can be specified for a request, not both.

optional before string: Return only items that occur before this marker. Before marker can be found from the response,
inside the 'paging' property. Note that only after or before can be specified for a request, not both.
*/
func (s *Session) GetClanMember(r ClanMemberRequest) ([]ClanMember, error) {
	member := ClanMemberList{}

	url, err := url.Parse(fmt.Sprintf("clans/%s/members", strings.Replace(r.ClanTag, "#", "%23", 1)))
	if err != nil {
		return member.Items, err
	}
	params := url.Query()

	if r.After > 0 {
		params.Add("after", fmt.Sprintf("%d", r.After))
	}
	if r.Before > 0 {
		params.Add("before", fmt.Sprintf("%d", r.Before))
	}
	if r.Limit > 0 {
		params.Add("limit", fmt.Sprintf("%d", r.Limit))
	}

	url.RawQuery = params.Encode()

	body, err := s.get(url.String())
	if err != nil {
		return member.Items, err
	}
	err = json.Unmarshal(body, &member)
	if err != nil {
		return member.Items, err
	}
	return member.Items, nil
}
