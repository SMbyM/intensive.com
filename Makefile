.PHONY: up down build logs

up:
	docker compose up --build

down:
	docker compose down

logs:
	docker compose logs -f backend
