// db/database.go
package db

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

// InitDB opens a connection to the SQLite database at the given filepath
// and verifies the connection. It sets a reasonable max open connections
// and returns any error encountered while opening or pinging the DB.
// filepath example: "../database/forum.db?_foreign_keys=1"
func InitDB(filepath string) error {
	var err error
	DB, err = sql.Open("sqlite3", filepath)
	if err != nil {
		return err
	}
	DB.SetMaxOpenConns(10)
	return DB.Ping()
}

// RunMigrations reads the SQL in `db/migrations.sql` and executes it
// against the configured database connection. Returns any execution error.
func RunMigrations() error {
	data, err := os.ReadFile("db/migrations.sql")
	if err != nil {
		return err
	}
	_, err = DB.Exec(string(data))
	return err
}
