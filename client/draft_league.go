package client

import (
	"github.com/matt-condon/fpl-draft-league-table-aggregator/core/models"
)

const (
	draftLeagueAddress = "https://draft.premierleague.com/api/league/"
)

func (c *Client) GetDraftLeague(leagueID string) (*models.League, error) {

	url := draftLeagueAddress + leagueID + "/details"

	request, err := c.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	league := &models.League{}

	_, err = c.Do(request, &league)
	if err != nil {
		return nil, err
	}

	return league, nil
}
