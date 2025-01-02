package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/matt-condon/fpl-draft-league-table-aggregator/client"
	"github.com/matt-condon/fpl-draft-league-table-aggregator/core/features"
	"github.com/matt-condon/fpl-draft-league-table-aggregator/core/models"
)

const (
	stageOneStaticLeagueRootPath = "data/draft-stage-one-league-table.json"
	stageOneJsonPath             = "data/view/stage-one-league-table.json"
	aggregatedJsonPath           = "data/view/aggregated-league-table-%d.json"
	aggregatedJsonPathLive       = "data/view/aggregated-league-table-live.json"
)

func main() {
	c := client.NewClient(nil)

	eventStatus := getEventStatus(c)
	currentEvent := eventStatus.Status[0].Event

	stageOneTable := getStageOneTable()
	stageTwoTable := getStageTwoTable(c, eventStatus.Status[0].Event)

	aggregatedTable := features.AggregateAndSort(*stageOneTable, *stageTwoTable, currentEvent)

	fmt.Println("\nStage 1 table:")
	stageOneTable.Display()

	fmt.Println("\nStage 2 table:")
	stageTwoTable.Display()

	fmt.Println("\nAggregated table:")
	aggregatedTable.Display()

	err := saveLeagueTableToJSON(*stageOneTable, stageOneJsonPath)
	if err != nil {
		fmt.Println("Error saving JSON:", err)
	}

	err = saveLeagueTableToJSON(aggregatedTable, fmt.Sprintf(aggregatedJsonPath, currentEvent))
	if err != nil {
		fmt.Println("Error saving JSON:", err)
	}

	err = saveLeagueTableToJSON(aggregatedTable, aggregatedJsonPathLive)
	if err != nil {
		fmt.Println("Error saving JSON:", err)
	}
}

func getEventStatus(c *client.Client) *models.EventStatusResponse {
	eventStatus, err := c.GetEventStatus()

	if err != nil {
		fmt.Println("Error retrieving event status", err)
		os.Exit(1)
	}
	return eventStatus
}

func getStageOneTable() *models.OrderedStandings {
	file, err := os.Open(stageOneStaticLeagueRootPath)
	if err != nil {
		log.Fatalf("Failed to open JSON file: %v", err)
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Failed to read JSON file: %v", err)
	}

	var root models.DraftRoot
	err = json.Unmarshal(bytes, &root)
	if err != nil {
		log.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	staticFinalGw := 19
	return features.NewOrderedStandings(root.LeagueEntries, root.Standings, staticFinalGw)
}

func getStageTwoTable(c *client.Client, eventID int) *models.OrderedStandings {
	leagueRoot, err := c.GetDraftLeague("186651")

	if err != nil {
		fmt.Println("Error retrieving draft league", err)
		os.Exit(1)
	}

	return features.NewOrderedStandings(leagueRoot.LeagueEntries, leagueRoot.Standings, eventID)
}

func saveLeagueTableToJSON(data models.OrderedStandings, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // Pretty print the JSON
	return encoder.Encode(data)
}
