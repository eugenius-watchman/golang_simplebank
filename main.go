package main

import (
	"database/sql"
	"log"

	"github.com/eugenius-watchman/golang_simplebank/api"
	db "github.com/eugenius-watchman/golang_simplebank/db/sqlc"
	"github.com/eugenius-watchman/golang_simplebank/util"
	_ "github.com/lib/pq"
	// "github.com/spf13/viper"
)

// Database connection constants
// const (
// 	dbDriver = "postgres" // PostgreSQL driver name
// 	// Connection string format: postgresql://user:password@host:port/database?params
// 	dbSource = "postgresql://root:secret@localhost:5433/simple_bank?sslmode=disable"

// 	serverAddress = "0.0.0.0:8080"
// )

func main() {
	// Load configuration
	config, err := util.LaodConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	// Establish database connection
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		// Fatal exits if connection fails
		log.Fatal("cannot connect to database:", err)
	}

	store := db.NewStore(conn)
	// server := api.NewServer(store)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server", err)
	}

}
