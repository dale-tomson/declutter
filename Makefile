APP_NAME := declutter
CMD_PATH := ./cmd/declutter
DIST_DIR := dist
VERSION_FILE := internal/version/version.go

# Get version from version.go
VERSION := $(shell grep 'const Version' $(VERSION_FILE) | sed 's/.*"\(.*\)"/\1/')

.PHONY: all build run test clean bump deps help

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Targets:'
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2}'

run: ## Run the application locally
	go run $(CMD_PATH)

build: ## Build for the current OS/Arch
	@echo "Building $(APP_NAME) v$(VERSION)..."
	go build -ldflags="-s -w" -o $(APP_NAME) $(CMD_PATH)

test: ## Run tests
	go test ./... -v

clean: ## Remove build artifacts
	rm -rf $(DIST_DIR)
	rm -f $(APP_NAME) $(APP_NAME).exe

bump: ## Bump version (Usage: make bump version=patch|minor|major|1.2.3)
	@if [ -z "$(version)" ]; then \
		./scripts/bump.sh; \
	else \
		./scripts/bump.sh $(version); \
	fi

deps: ## Install Go dependencies
	go mod download
	go install fyne.io/fyne/v2/cmd/fyne@latest

build-linux: ## Build for Linux (amd64)
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o $(DIST_DIR)/$(APP_NAME)-linux-amd64 $(CMD_PATH)

build-windows: ## Build for Windows (amd64)
	GOOS=windows GOARCH=amd64 go build -ldflags="-s -w -H windowsgui" -o $(DIST_DIR)/$(APP_NAME).exe $(CMD_PATH)
