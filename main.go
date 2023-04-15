package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"github.com/stinekamau/simplebank/api"
	db "github.com/stinekamau/simplebank/db/sqlc"
	"github.com/stinekamau/simplebank/utils"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		panic(fmt.Sprintf("couldn't load environment variables: %v", err))
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatalf("Error starting the database, %+v", err)
	}
	if err := conn.Ping(); err != nil {
		fmt.Printf("Database ping results to an error")
	}

	store := db.NewStore(conn)
	server := api.NewServer(config, store)
	server.Start(config.ServerAddress)
}
