package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	api "tutorial.sqlc.dev/app/api"
	db "tutorial.sqlc.dev/app/db/sqlc"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:secret@localhost:5432/simple_bank2?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {

	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to the database")
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.StartServer(serverAddress)

	if err != nil {
		log.Fatal("can not start server", err)
	}

}
