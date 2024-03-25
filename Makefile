# Run in docker-compose
up:
	docker compose up -d --force-recreate

up-dynamo:
	docker compose up -d --force-recreate dynamo-config

down:
	docker compose down --remove-orphans --rmi local

# Build and run locally
deps:
	go mod tidy
	go mod vendor

run:
	go run main.go -e 