package GoClans

import "encoding/json"

//Get information about the current gold pass season.
func (s *Session) GetCurrentGoldpassSeason() (GoldPassSeason, error) {
	Season := GoldPassSeason{}
	body, err := s.get("goldpass/seasons/current")
	if err != nil {
		return Season, err
	}
	err = json.Unmarshal(body, &Season)
	if err != nil {
		return Season, err
	}
	return Season, nil
}
