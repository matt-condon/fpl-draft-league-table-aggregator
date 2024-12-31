package draft

import (
	"sort"

	"github.com/matt-condon/fpl-draft-league-table-aggregator/core/models"
)

// Aggregate and sort the standings based on PlayerName
func AggregateAndSort(standings1, standings2 models.OrderedStandings) models.OrderedStandings {
	// Map to aggregate totals by PlayerName
	aggregated := make(map[string]*models.OrderedStanding)

	// Iterate through the first slice and store the standings
	for _, standing := range standings1.Standings {
		aggregated[standing.PlayerName] = &models.OrderedStanding{
			Rank:       standing.Rank,
			EntryName:  standing.EntryName,
			PlayerName: standing.PlayerName,
			EventTotal: standing.EventTotal,
			Total:      standing.Total,
		}
	}

	// Iterate through the second slice, updating the map
	for _, standing := range standings2.Standings {
		// If the player already exists, update the Total, else add a new entry
		if entry, exists := aggregated[standing.PlayerName]; exists {
			entry.Total += standing.Total
			entry.Rank = standing.Rank
			entry.EntryName = standing.EntryName
			entry.EventTotal = standing.EventTotal
		} else {
			// Player is not in the first slice, so directly add the entry
			aggregated[standing.PlayerName] = &models.OrderedStanding{
				Rank:       standing.Rank,
				EntryName:  standing.EntryName,
				PlayerName: standing.PlayerName,
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

	// Return the sorted and aggregated standings
	return models.OrderedStandings{Standings: result}
}
