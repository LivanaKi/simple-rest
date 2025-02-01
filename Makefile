build:
		go build -v ./
.PHONY: build

migrate-up:
	migrate -path ./migration -database 'postgres://root:root@localhost:5433/rest?sslmode=disable' up
.PHONY: migrate-up

migrate-down:
	migrate -path ./migration -database 'postgres://root:root@localhost:5433/rest?sslmode=disable' down

test:
	go test ./controller ./service ./repository

test-brench:
	go test -bench . ./concatenations

lint: ## Run GoLangCI Lint
	golangci-lint run ./... --config ./.golangci.yml

.DEFAULT_GOAL := build