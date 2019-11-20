.DEFAULT_GOAL := help

test: lint ## run go test ## make test
	go test -race -v ./...

lint: ## run all lint ## make lint
	@golangci-lint run ./...
