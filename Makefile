.PHONY: default
default: up

HAS_LINT := $(shell command -v golangci-lint;)

bootstrap:
ifndef HAS_LINT
	go get -u github.com/golangci/golangci-lint/cmd/golangci-lint
endif
ifndef HAS_IMPORTS
	go get -u golang.org/x/tools/cmd/goimports
endif

.PHONY: lint
lint: ## Run linter for all project files
	@echo "Running linter..."
	@golangci-lint run
	@echo "Done"

.PHONY: test
test: ## Run tests
	@echo "Running tests..."
	bash -c "go clean -testcache"
	bash -c "go test ./... -v"
	@echo "Done"

.PHONY: build
build: ## Run build project
	bash -c "docker-compose -f deployments/docker-compose.yml build"

.PHONY: up
up: ## Run build and up project
	bash -c "docker-compose -f deployments/docker-compose.yml up --build"

.PHONY: start
start: ## Start project
	bash -c "docker-compose -f deployments/docker-compose.yml start"

.PHONY: stop
stop: ## Stop project
	bash -c "docker-compose -f deployments/docker-compose.yml stop"

.PHONY: poker
poker: ## Run parsing task
	bash -c "go run cmd/poker/main.go"

.PHONY: client
client: ## Run parsing task
	bash -c "go run cmd/client/main.go"