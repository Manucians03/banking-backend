include app.env

docker_network:
	docker network create banking

docker_postgres:
	docker run --name postgres -p $(DOCKER_DB_PORT):$(DB_PORT) -e POSTGRES_USER=$(DB_USERNAME) -e POSTGRES_PASSWORD=$(DB_PASSWORD) -d postgres:16.3-alpine3.20

docker_build:
	docker build -t banking:latest .

docker_run:
	docker run --name banking --network banking -p 8080:8080 -e DB_SOURCE=$(DB_SOURCE_DOCKER) -d banking:latest

createdb:
	docker exec -it postgres createdb --username=$(DB_USERNAME) --owner=$(DB_USERNAME) $(DB_NAME)

dropdb:
	docker exec -it postgres dropdb $(DB_NAME)

migrateup:
	migrate -path db/migration -database "$(DB_SOURCE)" -verbose up

migratedown:
	migrate -path db/migration -database "$(DB_SOURCE)" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -destination db/mock/store.go github.com/Manucians03/banking-backend/db/sqlc Store  

.PHONY: docker_network docker_postgres docker_build docker_run createdb dropdb migrateup migratedown sqlc test server mock