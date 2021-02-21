package clients

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
)
var (
	sqliteClient *sql.DB
)
func init() {
	// Set client options
	sqliteClient, _ = sql.Open("sqlite3", "./database/sample.db") // Open the created SQLite File
}

func GetSQLClient() *sql.DB {
	return sqliteClient
}
