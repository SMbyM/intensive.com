include .env.example
export


.PHONY: up down build logs

up:
	sqlc generate
	docker compose up --build

down:
	docker compose down

logs:
	docker compose logs -f backend


.PHONY: db-up db-down migrate-up migrate-down migrate-create

db-up:
	docker compose up -d db

db-down:
	docker compose stop db

migrate-up: db-up
	migrate -path migrations -database "$(DATABASE_URL_LOCAL)" up

migrate-down:
	migrate -path migrations -database "$(DATABASE_URL_LOCAL)" down 1

migrate-create:
	migrate create -ext sql -dir migrations -seq $(name)