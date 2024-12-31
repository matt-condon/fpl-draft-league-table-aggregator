.PHONY: draft-league
draft-league: ## Display the draft league
	@echo "==> Displaying draft league"
	go run core/workflows/aggregate_draft_league/main.go
