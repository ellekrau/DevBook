package database

import (
	"api/src/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // MySQL driver
)

// Connect open and return a DB connection
func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", config.DbConnectionString)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
