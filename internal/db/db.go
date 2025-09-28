package db

import (
	"database/sql"
	"tessa/internal/config"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// Structure table for storing data
type Clipboard struct {
	id			uint
	data		string
	updated_at	time.Time
}

var DB *sql.DB

// Create a database connection to the file created during the configuration of the applicarion
func GetCon() {
	path := config.GetDBPath()
	DB, _ = sql.Open("sqlite3", path)
}

// Create the tables for storing the data in the database file
func Init() error {
	sqlStmt := `
		CREATE TABLE IF NOT EXISTS tessa (
			id			INTEGER	PRIMARY KEY	AUTOINCREMENT,
			data		TEXT NOT NULL,
			updated_at	DATETIME DEFAULT CURRENT_TIMESTAMP
		);
	`

	if _, err := DB.Exec(sqlStmt); err != nil {
		return err
	}

	return nil	
}

// Closes any existing database connection (must be called with 'defer' keyword)
func CloseCon() {
	DB.Close()
}