.DEFAULT_GOAL := help

deps: ## Install dependencies ## make deps
	@export GO111MODULE=off \
	  go get github.com/volatiletech/sqlboiler \
	  go get github.com/volatiletech/sqlboiler/drivers/sqlboiler-mysql \
	  go get github.com/rubenv/sql-migrate/...

test: lint ## run go test ## make test
	go test -v ./...

lint: lint_go ## run all lint ## make lint

lint_go: ## lint all go files ## make lint_go
	@golangci-lint run ./...
