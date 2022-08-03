package GoClans

import (
	"encoding/json"
	"strings"
)

/*
Get information about a single player by player tag. Player tags can be found either in game or by from clan member lists.

parameters:

required playerTag string: Tag of the player
*/
func (s *Session) GetPlayer(tag string) (Player, error) {
	player := Player{}
	body, err := s.get("players/" + strings.Replace(tag, "#", "%23", 1))
	if err != nil {
		return player, err
	}
	err = json.Unmarshal(body, &player)
	if err != nil {
		return player, err
	}
	return player, nil
}

/*
Verify player API token that can be found from the game settings. This API call can be used to check that players
own the game accounts they claim to own as they need to provide the one-time use API token that exists inside the game.

parameters:

required playerTag string: Tag of the Player

required token string: token found in the game
*/
func (s *Session) VerifyPlayer(tag, token string) (VerifyTokenResponse, error) {
	verifyRespone := VerifyTokenResponse{}
	body, err := s.post("players/"+strings.Replace(tag, "#", "%23", 1)+"/verifytoken", VerifyTokenRequest{Token: token})
	if err != nil {
		return verifyRespone, err
	}
	err = json.Unmarshal(body, &verifyRespone)
	if err != nil {
		return verifyRespone, err
	}
	return verifyRespone, nil
}
