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
	go mod download && go mod verify
	@echo "Dependencies installed $(OK_STRING)"

run-db:
	@echo "Starting database..."
	docker compose up -d
	@echo "Database running $(OK_STRING)"

run-api:
	@echo "Starting API..."
	go run main.go
	@echo "API running $(OK_STRING)"

build-docker:
	@echo "Building in docker..."
	docker build -t simple-crud-api .
	@echo "Application has been built in docker $(OK_STRING)"

run-docker:
	@echo "Starting in docker..."
	docker run -dp 3000:3000 simple-crud-api
	@echo "Application running in docker $(OK_STRING)"
