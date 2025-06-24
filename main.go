package main

import (
	"database/sql"
	"log"

	"github.com/eugenius-watchman/golang_simplebank/api"
	db "github.com/eugenius-watchman/golang_simplebank/db/sqlc"
	_ "github.com/lib/pq"
)

// Database connection constants
const (
	dbDriver = "postgres" // PostgreSQL driver name
	// Connection string format: postgresql://user:password@host:port/database?params
	dbSource = "postgresql://root:secret@localhost:5433/simple_bank?sslmode=disable"

	serverAddress = "0.0.0.0:8080"
)

func main() {
	// Establish database connection
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		// Fatal exits if connection fails
		log.Fatal("Cannot connect to database:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("Cannot start server", err)
	}

}
