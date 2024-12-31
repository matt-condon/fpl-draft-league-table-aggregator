package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/matt-condon/fpl-draft-league-table-aggregator/core/features/draft"
	"github.com/matt-condon/fpl-draft-league-table-aggregator/core/models"
)

func main() {
	stageOneTable := getStageOneTable()
	stageTwoTable := getStageTwoTable()
	aggregatedTable := draft.AggregateAndSort(*stageOneTable, *stageTwoTable)

	fmt.Println("\nStage 1 table:")
	stageOneTable.Display()

	fmt.Println("\nStage 2 table:")
	stageTwoTable.Display()

	fmt.Println("\nAggregated table:")
	aggregatedTable.Display()

	err := saveLeagueTableToJSON(aggregatedTable, "data/aggregated-league-table.json")
	if err != nil {
		fmt.Println("Error saving JSON:", err)
	}
}

func getStageOneTable() *models.OrderedStandings {
	jsonFilePath := "data/draft-stage-one-league-table-partial.json"

	file, err := os.Open(jsonFilePath)
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

	return draft.NewOrderedStandings(root.LeagueEntries, root.Standings)
}

func getStageTwoTable() *models.OrderedStandings {
	jsonFilePath := "data/draft-stage-two-league-table-temp.json"

	file, err := os.Open(jsonFilePath)
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

	return draft.NewOrderedStandings(root.LeagueEntries, root.Standings)
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
