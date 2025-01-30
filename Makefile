build:
		go build -v ./
.PHONY: build

migrate-up:
	migrate -path ./migration -database 'postgres://root:root@localhost:5433/rest?sslmode=disable' up
.PHONY: migrate-up

migrate-down:
	migrate -path ./migration -database 'postgres://root:root@localhost:5433/rest?sslmode=disable' down
.PHONY: migrate-down

.DEFAULT_GOAL := build