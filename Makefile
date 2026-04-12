include .env
export

export PROJECT_ROOT=$(shell pwd)

env-up:
	docker compose up -d

env-down:
	docker compose down

env-cleanup:
	@read -p "Delete all volume env files? [y/N]: " ans; \
	if [ "$$ans" = "y" ]; then \
		docker compose down && \
		docker volume rm tasks_postgres_data 2>/dev/null || true && \
		echo "env files deleted"; \
	else \
		echo "canceled"; \
	fi

env-port-forward:
	@docker compose up -d port-forwarder

env-port-close:
	@docker compose down port-forwarder

migrate-create:
	@if [ -z "$(seq)" ]; then \
		echo "seq param required"; \
		exit 1; \
	fi; \
	docker compose run --rm tasks-migrate \
		create \
		-ext sql \
		-dir /migrations \
		-seq "$(seq)"

migrate-up:
	make migrate-action action=up

migrate-down:
	make migrate-action action=down

migrate-action:
	docker compose --env-file .env run --rm tasks-migrate \
		-path /migrations \
		-database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:5432/${POSTGRES_DB}?sslmode=disable \
		"$(action)"

build:
	docker compose build

build-app:
	docker compose build tasks-app

run:
	docker compose build && docker compose up

run-db-only:
	docker compose up tasks-postgres port-forwarder