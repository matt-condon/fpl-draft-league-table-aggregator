package models

type EventStatus struct {
	BonusAdded     bool   `json:"bonus_added"`
	Date           string `json:"date"`
	Event          int    `json:"event"`
	LeaguesUpdated bool   `json:"leagues_updated"`
	Points         string `json:"points"`
}

type EventStatusResponse struct {
	Status  []EventStatus `json:"status"`
	Leagues string        `json:"leagues"`
}
