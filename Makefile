# Purpose: Defines common commands for building, running, testing, and deploying your application.
# Examples: make build, make run, make test, make migrate-up.
# Why: Standardizes development workflows and makes it easy for new contributors to get started.


run: ## Run the application
	
	go run cmd/api/main.go

build: ## Build the application
	
	go cmd/api/build