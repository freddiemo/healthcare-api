build:
	docker compose -f docker-compose.yml build --force-rm
run:
	docker compose -f docker-compose.yml up
build-tests:
	docker compose -f docker-compose.test.yml build --force-rm
run-tests:
	docker compose -f docker-compose.test.yml up --abort-on-container-exit
	docker compose -f docker-compose.test.yml down
