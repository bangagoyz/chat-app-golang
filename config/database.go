package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	db *sql.DB
}

func DatabaseConnection() (*Database, error) {
	sqlInfo := "username:password@tcp(localhost:3306)/go_chat"

	db, err := sql.Open("mysql", sqlInfo)
	if err != nil {
		log.Println("Database connection error", err)
	}
	err = db.Ping()
	if err != nil {
		log.Println("Failed to ping database", err)
	}
	log.Println("Database connected")
	return &Database{db: db}, nil
}

func (d *Database) Close() {
	d.db.Close()
}

func (d *Database) GetDB() *sql.DB {
	return d.db
}
