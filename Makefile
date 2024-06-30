postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=minhdeptrai123 -d postgres:16.3-alpine3.20

createdb:
	docker exec -it postgres createdb --username=root --owner=root banking

dropdb:
	docker exec -it postgres dropdb banking

migrateup:
	migrate -path db/migration -database "postgresql://root:minhdeptrai123@localhost:5432/banking?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:minhdeptrai123@localhost:5432/banking?sslmode=disable" -verbose down

.PHONY: postgres createdb dropdb migrateup migratedown