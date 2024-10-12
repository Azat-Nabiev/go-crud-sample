APP_NAME := myapp
MIGRATION_NAME := migrations
BUILD_DIR := build

build-api:
	go build -o $(BUILD_DIR)/$(APP_NAME) ./cmd/awesomeProject/main.go

build-migrations:
	go build -o $(BUILD_DIR)/$(MIGRATION_NAME) ./cmd/migrate/main.go

build: build-api build-migrations

run: build-api
	@echo "Starting api..."
	@bash -c "set -a && source .env && set +a && ./$(BUILD_DIR)/$(APP_NAME)"

migrate: build-migrations
	@echo "Add migrations..."
	@bash -c "set -a && source .env && set +a && ./$(BUILD_DIR)/$(MIGRATION_NAME)"

.PHONY: build build-api build-migrations run migrate



