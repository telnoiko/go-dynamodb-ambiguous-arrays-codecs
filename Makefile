# Run in docker-compose
up:
	docker compose up -d --force-recreate

up-dynamo:
	docker compose up -d --force-recreate dynamodb-local dynamo-config

down:
	docker compose down --remove-orphans --rmi local

# Build and run locally
deps:
	go mod tidy
	go mod vendor

run:
	DYANMODB_HOST=http://localhost:8000 go run main.go