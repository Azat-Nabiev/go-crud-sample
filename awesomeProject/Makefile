APP_NAME := myapp
MIGRATION_NAME := migrations
BUILD_DIR := build
COMPOSE_FILE=docker-compose.yml

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

docker-up:
	docker-compose -f $(COMPOSE_FILE) up -d

docker-down:
	docker-compose -f $(COMPOSE_FILE) down

docker-restart: docker-down docker-up

docker-logs:
	docker-compose -f $(COMPOSE_FILE) logs

.PHONY: build build-api build-migrations run migrate docker-up docker-down docker-logs docker-restart



