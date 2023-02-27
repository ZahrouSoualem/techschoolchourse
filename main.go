package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	api "tutorial.sqlc.dev/app/api"
	db "tutorial.sqlc.dev/app/db/sqlc"
	"tutorial.sqlc.dev/app/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot get to the file")
	}
	fmt.Println(config.DBDriver)
	fmt.Println(config.ServerAddress)
	fmt.Println(config.DBSource)
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to the database")
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.StartServer(config.ServerAddress)

	if err != nil {
		log.Fatal("can not start server", err)
	}

}
