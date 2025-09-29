package db

import (
	"database/sql"
	"tessa/internal/config"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// Structure table for storing data
type Clipboard struct {
	Id			uint
	Data		string
	UpdatedAt	time.Time
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

// Save data to database
func SaveData(message string) error {
	sqlStmt := `
		INSERT INTO tessa ( data ) VALUES ( $1 );
	`
	if _, err := DB.Exec(sqlStmt, message); err != nil {
		return err
	}
	return nil
}

func GetData() ([]Clipboard, error) {
	sqlStmt := `
		SELECT id, data, updated_at FROM tessa;
	`

	var clipData []Clipboard

	row, err := DB.Query(sqlStmt); 
	if err != nil {
		return nil, err
	}

	defer row.Close()
	for row.Next() {
		var clip Clipboard
		if err := row.Scan(&clip.Id, &clip.Data, &clip.UpdatedAt); err != nil {
			return nil, err
		}

		clipData = append(clipData, clip);
	}

	if err := row.Err(); err != nil {
		return nil, err
	}

	return clipData, nil
}

func CleanData() error {
	sqlStmt := `
		DELETE FROM tessa
	`

	if _, err := DB.Exec(sqlStmt); err != nil {
		return err
	}

	return nil
}

func GetDataById(id int) (Clipboard, error) {
	sqlStml := `SELECT id, data FROM tessa WHERE id = $1;`

	var row Clipboard
	if err := DB.QueryRow(sqlStml, id).Scan(
		&row.Id,
		&row.Data,
	); err != nil {
		return Clipboard{}, err
	}

	return row, nil
}

// Closes any existing database connection (must be called with 'defer' keyword)
func CloseCon() {
	DB.Close()
}