package utils

import (
	"database/sql"
	"os"
	"time"
)

func DbConnect() (*sql.DB, error) {
	dbURI := os.Getenv("DB_URI")
	db, err := sql.Open("mysql", dbURI)
	if err != nil {
		Log("ERROR", "db", "unable to connect to db because %v", err)
		return nil, err
	}

	if err = db.Ping(); err != nil {
		Log("ERROR", "db", "unable to ping db because %v", err)
		return nil, err
	}

	db.SetMaxIdleConns(64)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(10 * time.Second)
	Log("INFO", "db", "Connected")
	return db, nil
}
