init:
	docker compose up -d --build

build:
	docker compose build

up:
	docker compose up -d

shell:
	docker compose exec app /bin/sh

down:
	docker compose down

sql:
	docker compose exec db psql -U postgres -d echo
