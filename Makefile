NO_COLOR=\x1b[0m
OK_COLOR=\x1b[32;01m
ERROR_COLOR=\x1b[31;01m
WARN_COLOR=\x1b[33;01m
OK_STRING=$(OK_COLOR)[OK]$(NO_COLOR)
ERROR_STRING=$(ERROR_COLOR)[ERRORS]$(NO_COLOR)
WARN_STRING=$(WARN_COLOR)[WARNINGS]$(NO_COLOR)

default: install

install:
	@echo "Installing dependencies..."
	go mod download

db:
	@echo "Starting database $(OK_STRING)"
	docker compose up -d

api:
	@echo "Starting API..."
	go run main.go

dockerize:
	@echo "Starting dockerize $(OK_STRING)"
	docker build -t simple-crud-api .
