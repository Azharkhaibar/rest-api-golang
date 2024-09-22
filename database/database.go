package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func InitDb() (*sql.DB, error) {
	dsn := "root@tcp(localhost:3306)/todos_golang"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// Check if the connection is successful
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	log.Println("Database connection established")
	return db, nil
}
