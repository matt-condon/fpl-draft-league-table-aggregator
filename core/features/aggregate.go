package features

import (
	"sort"

	"github.com/matt-condon/fpl-draft-league-table-aggregator/core/models"
)

// Aggregate and sort the standings based on PlayerName
func AggregateAndSort(standings1, standings2 models.OrderedStandings, eventID int) models.OrderedStandings {
	// Map to aggregate totals by PlayerName
	aggregated := make(map[string]*models.OrderedStanding)

	// Iterate through the first slice and store the standings
	for _, standing := range standings1.Standings {
		aggregated[standing.PlayerName] = &models.OrderedStanding{
			EntryName:  standing.EntryName,
			PlayerName: standing.PlayerName,
			TeamUrl:    standing.TeamUrl,
			EventTotal: standing.EventTotal,
			Total:      standing.Total,
		}
	}

	// Iterate through the second slice, updating the map
	for _, standing := range standings2.Standings {
		// If the player already exists, update the Total, else add a new entry
		if entry, exists := aggregated[standing.PlayerName]; exists {
			entry.EntryName = standing.EntryName
			entry.TeamUrl = standing.TeamUrl
			entry.EventTotal = standing.EventTotal
			entry.StageTwoTotal = standing.Total
			entry.Total += standing.Total
		} else {
			// Player is not in the first slice, so directly add the entry
			aggregated[standing.PlayerName] = &models.OrderedStanding{
				EntryName:  standing.EntryName,
				PlayerName: standing.PlayerName,
				TeamUrl:    standing.TeamUrl,
				EventTotal: standing.EventTotal,
				Total:      standing.Total,
			}
		}
	}

	// Convert the aggregated map back to a slice
	var result []models.OrderedStanding
	for _, standing := range aggregated {
		result = append(result, *standing)
	}

	// Sort the result slice by Total in descending order
	sort.Slice(result, func(i, j int) bool {
		return result[i].Total > result[j].Total
	})

	// Assign ranks based on sorted order
	for i := range result {
		result[i].Rank = i + 1 // Rank starts from 1
	}

	// Return the sorted and aggregated standings
	return models.OrderedStandings{Standings: result, Event: eventID}
}
