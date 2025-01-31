.DEFAULT_GOAL := help

.DEFAULT_GOAL := help

.PHONY: setup
setup: ## install linter
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.45.2
	
.PHONY: test
test: ## Run tests.
	go test -race -cover -covermode=atomic -coverprofile=coverage.out ./...

.PHONY: lint
lint: ## Run golangci-lint.
	golangci-lint run -v --timeout=5m

.PHONY: help
help:
	@grep -E '^[%a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
