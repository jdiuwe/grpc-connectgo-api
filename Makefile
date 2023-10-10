.PHONY: lint

generate-db-models:
	sqlc generate

migrate-up:
	 migrate --path db/migrations --database "postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable" --verbose up

migrate-down:
	 migrate --path db/migrations --database "postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable" --verbose down

lint:
	golangci-lint run

docker-up:
	docker-compose -f docker-compose.yml up -d
