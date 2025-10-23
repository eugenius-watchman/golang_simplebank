// Package db contains database access logic and test utilities
package db

import (
	"database/sql" // Standard Go SQL package
	"log"          // For logging errors
	"os"           // For OS interaction (test exit codes)
	"testing"      // Go testing framework

	// PostgreSQL driver (imported for side-effects only)
	// The underscore indicates we're importing it solely for its initialization
	"github.com/eugenius-watchman/golang_simplebank/util"
	_ "github.com/lib/pq"
)

// Database connection constants
// const (
// 	dbDriver = "postgres"  // PostgreSQL driver name
// 	// Connection string format: postgresql://user:password@host:port/database?params
// 	dbSource = "postgresql://root:secret@localhost:5433/simple_bank?sslmode=disable"
// )

// Global variables shared across tests
var (
	testQueries *Queries // Generated SQL queries (from sqlc)
	testDB      *sql.DB  // Database connection pool
)

// TestMain is the entry point for database tests
// It runs before any other tests in the package
func TestMain(m *testing.M) {
	// Load config from root directory (where app.env lives)
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	// Establish database connection using config
	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to database:", err)
	}
	defer testDB.Close()

	// Initialize queries
	testQueries = New(testDB)

	// Run tests
	os.Exit(m.Run())

	// VerifyGoVersion("go1.22.0")

	// var err error

	// // Establish database connection
	// testDB, err = sql.Open(dbDriver, dbSource)
	// if err != nil {
	// 	// Fatal exits if connection fails
	// 	log.Fatal("Cannot connect to database:", err)
	// }
	// // Ensure connection closes when tests complete
	// defer testDB.Close()

	// // Initialize query interface with our connection
	// testQueries = New(testDB)

	// // Run all tests and exit with proper status code
	// os.Exit(m.Run())
}
