build:
		go build -o simple-rest ./cmd/main.go

migrate-up:
	migrate -path ./migration -database 'postgres://root:root@localhost:5433/rest?sslmode=disable' up

migrate-down:
	migrate -path ./migration -database 'postgres://root:root@localhost:5433/rest?sslmode=disable' down

test:
	go test ./internal/controller ./internal/service ./internal/repository

test-bench:
	go test -bench . ./concatenations

lint:
	golangci-lint run ./... --config ./.golangci.yml

.DEFAULT_GOAL := build