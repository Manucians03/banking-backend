include app.env

postgres:
	docker run --name postgres -p $(DB_DOCKER_PORT):$(DB_PORT) -e POSTGRES_USER=$(DB_USERNAME) -e POSTGRES_PASSWORD=$(DB_PASSWORD) -d postgres:16.3-alpine3.20

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

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server mock