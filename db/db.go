package db

import (
	"database/sql"
	"os"
	"path"

	_ "github.com/mattn/go-sqlite3"
)

const DB_FILE = "hyper.db"

type Database struct {
	instance *sql.DB
	filepath string
}

// Returns new instance of Database
func New(rootDir string) *Database {
	var err error

	if len(rootDir) == 0 {
		rootDir, err = os.UserHomeDir()

		if err != nil {
			rootDir = "."
		}
	}

	return &Database{
		instance: nil,
		filepath: path.Join(rootDir, DB_FILE),
	}
}

// Returns file path of selected database file
func (d *Database) FilePath() string {
	return d.filepath
}

// Opens new database connection
func (d *Database) Open() (err error) {
	if d.instance != nil {
		return nil
	}

	d.instance, err = sql.Open("sqlite3", d.filepath)
	return err
}

// Closes active database connection
func (d *Database) Close() (err error) {
	if d.instance == nil {
		return nil
	} else if err = d.instance.Close(); err != nil {
		d.instance = nil
		return nil
	} else {
		return err
	}
}
