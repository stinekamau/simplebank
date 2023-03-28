package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"github.com/stinekamau/simplebank/api"
	db "github.com/stinekamau/simplebank/db/sqlc"
)

func main() {
	const (
		dbDriver = "postgres"
		dbSource = "postgresql://root:secret@localhost:5432/bank?sslmode=disable"
	)

	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatalf("Error starting the database, %+v", err)
	}
	if err := conn.Ping();err!=nil{
		fmt.Printf("Database ping results to an error")
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)
	server.Start(":8080")
}
