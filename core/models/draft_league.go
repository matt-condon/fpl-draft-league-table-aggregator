package models

import (
	"fmt"
	"time"
)

// DraftRoot represents the entire JSON structure
type DraftRoot struct {
	League        League        `json:"league"`
	LeagueEntries []LeagueEntry `json:"league_entries"`
	Standings     []Standing    `json:"standings"`
}

// League represents the league information
type League struct {
	AdminEntry           int       `json:"admin_entry"`
	Closed               bool      `json:"closed"`
	DraftDateTime        time.Time `json:"draft_dt"`
	DraftPickTimeLimit   int       `json:"draft_pick_time_limit"`
	DraftStatus          string    `json:"draft_status"`
	DraftTimezoneDisplay string    `json:"draft_tz_show"`
	ID                   int       `json:"id"`
	KORounds             int       `json:"ko_rounds"`
	MakeCodePublic       bool      `json:"make_code_public"`
	MaxEntries           int       `json:"max_entries"`
	MinEntries           int       `json:"min_entries"`
	Name                 string    `json:"name"`
	Scoring              string    `json:"scoring"`
	StartEvent           int       `json:"start_event"`
	StopEvent            int       `json:"stop_event"`
	Trades               string    `json:"trades"`
	TransactionMode      string    `json:"transaction_mode"`
	Variety              string    `json:"variety"`
}

// LeagueEntry represents individual entries in the league
type LeagueEntry struct {
	EntryID         int       `json:"entry_id"`
	EntryName       string    `json:"entry_name"`
	ID              int       `json:"id"`
	JoinedTime      time.Time `json:"joined_time"`
	PlayerFirstName string    `json:"player_first_name"`
	PlayerLastName  string    `json:"player_last_name"`
	ShortName       string    `json:"short_name"`
	WaiverPick      int       `json:"waiver_pick"`
}

// Standing represents standings in the league
type Standing struct {
	EventTotal  int `json:"event_total"`
	LastRank    int `json:"last_rank"`
	LeagueEntry int `json:"league_entry"`
	Rank        int `json:"rank"`
	RankSort    int `json:"rank_sort"`
	Total       int `json:"total"`
}

// OrderedStanding represents a structured standing with entry details
type OrderedStanding struct {
	Rank          int
	EntryName     string
	TeamUrl       string
	PlayerName    string
	EventTotal    int
	StageTwoTotal int
	Total         int
}

// OrderedStandings contains a slice of OrderedStanding and methods to manipulate or display them
type OrderedStandings struct {
	Standings []OrderedStanding
	Event     int
}

// Display outputs the standings in a formatted table
func (os *OrderedStandings) Display() {
	fmt.Printf("%-5s %-20s %-20s %-12s %-18s %-12s\n", "Rank", "Entry Name", "Player Name", "Event Total", "Stage Two Total", "Total")
	for _, standing := range os.Standings {
		fmt.Printf("%-5d %-20s %-20s %-12d %-18d %-12d\n", standing.Rank, standing.EntryName, standing.PlayerName, standing.EventTotal, standing.StageTwoTotal, standing.Total)
	}
}
