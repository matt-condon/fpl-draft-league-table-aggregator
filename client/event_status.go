package client

import (
	"github.com/matt-condon/fpl-draft-league-table-aggregator/core/models"
)

const (
	eventStatusAddress = "https://draft.premierleague.com/api/pl/event-status"
)

func (c *Client) GetEventStatus() (*models.EventStatusResponse, error) {

	request, err := c.NewRequest("GET", eventStatusAddress, nil)
	if err != nil {
		return nil, err
	}

	eventStatus := &models.EventStatusResponse{}

	_, err = c.Do(request, &eventStatus)
	if err != nil {
		return nil, err
	}

	return eventStatus, nil
}
