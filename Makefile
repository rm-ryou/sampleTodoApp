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

.PHONY: mysql
mysql: up
	docker compose exec mysql mysql -u user -p

.PHONY: test
test: _test_api

.PHONY: _test_api
_test_api:
	@cd docker/go; go test -v ./...

