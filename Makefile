.PHONY: run
run: build up

.PHONY: build
build:
	docker compose build

.PHONY: up
up:
	docker compose up --detach

.PHONY: down
down:
	docker compose down

.PHONY: re
re: down run
