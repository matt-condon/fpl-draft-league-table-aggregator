package features

import (
	"fmt"
	"sort"

	"github.com/matt-condon/fpl-draft-league-table-aggregator/core/models"
)

const (
	teamUrlTemplate = "https://draft.premierleague.com/entry/%d/event/%d"
)

// NewOrderedStandings creates an OrderedStandings object from league entries and standings
func NewOrderedStandings(entries []models.LeagueEntry, standings []models.Standing, eventId int) *models.OrderedStandings {
	// Create a map for quick lookup of LeagueEntry by ID
	entryMap := make(map[int]models.LeagueEntry)
	for _, entry := range entries {
		entryMap[entry.ID] = entry
	}

	// Build the OrderedStanding slice
	orderedStandings := []models.OrderedStanding{}
	for _, standing := range standings {
		entry, exists := entryMap[standing.LeagueEntry]
		if !exists {
			fmt.Printf("entry not found: %-5d in ", entry.EntryID)
			continue // Skip if entry not found
		}
		orderedStandings = append(orderedStandings, models.OrderedStanding{
			Rank:       standing.Rank,
			EntryName:  entry.EntryName,
			TeamUrl:    fmt.Sprintf(teamUrlTemplate, entry.EntryID, eventId),
			PlayerName: entry.PlayerFirstName + " " + entry.PlayerLastName,
			EventTotal: standing.EventTotal,
			Total:      standing.Total,
		})
	}

	// Sort standings by rank
	sort.Slice(orderedStandings, func(i, j int) bool {
		return orderedStandings[i].Rank < orderedStandings[j].Rank
	})

	return &models.OrderedStandings{Standings: orderedStandings, Event: eventId}
}
