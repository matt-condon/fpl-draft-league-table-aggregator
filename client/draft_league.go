package client

import (
	"github.com/matt-condon/fpl-draft-league-table-aggregator/core/models"
)

const (
	draftLeagueAddress = "https://draft.premierleague.com/api/league/"
)

func (c *Client) GetDraftLeague(leagueID string) (*models.DraftRoot, error) {

	url := draftLeagueAddress + leagueID + "/details"

	request, err := c.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	root := &models.DraftRoot{}

	_, err = c.Do(request, &root)
	if err != nil {
		return nil, err
	}

	return root, nil
}
