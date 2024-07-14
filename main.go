package main

import (
	"database/sql"
	"log"

	"github.com/Manucians03/banking-backend/api"
	db "github.com/Manucians03/banking-backend/db/sqlc"
	_ "github.com/lib/pq"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:minhdeptrai123@localhost:5432/banking?sslmode=disable"
	serverAddress = "127.0.0.1:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Cannot connect to db: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("Cannot start server: ", err)
	}
}
