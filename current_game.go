package lol

import "fmt"

type CurrentGame struct {
	BannedChampion []struct {
		Champion ChampionID `json:"championId"`
		PickTurn int        `json:"pickTurn"`
		Team     TeamID     `json:"teamID"`
	} `json:"bannedChampions"`

	Id         GameID `json:"gameId"`
	GameLength int64  `json:"GameLength"`

	Participants []struct {
		Id   SummonerID `json:"summonerId`
		Name string     `json:"summonerName`
	} `json:"participants"`

	Observer struct {
		EncryptionKey string `json:"encryptionKey"`
	} `json:"observers"`
}

func (a *APIRegionalEndpoint) GetCurrentGame(id SummonerID) (*CurrentGame, error) {
	res := &CurrentGame{}
	err := a.g.Get(fmt.Sprintf("https://%s/observer-mode/rest/consumer/getSpectatorGameInfo/%s/%d?api_key=%s",
		a.region.url,
		a.region.platformId,
		id,
		a.key), res)
	if rerr, ok := err.(RESTError); ok == true {
		if rerr.Code == 404 {
			//user is just not in a game currently
			return nil, nil
		}
	}

	if err != nil {
		return res, fmt.Errorf("Could not get current game infor for %d: %s", id, err)
	}
	return res, nil
}
