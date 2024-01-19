package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitDatabase() {
	var err error
	db, err = sql.Open("mysql", "root:@tcp(localhost:3306)/")
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to the database")
}

func CreateDatabase() {
	_, err := db.Exec("CREATE DATABASE IF NOT EXISTS bookcollect")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database 'bookcollect' created successfully")
}

func UseDatabase() {
	_, err := db.Exec("USE bookcollect")
	if err != nil {
		log.Fatal(err)
	}
}

func CreateTable() {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS books (
			id INT AUTO_INCREMENT PRIMARY KEY,
			title VARCHAR(255),
			author VARCHAR(255),
			rating INT,
			created_at TIMESTAMP,
			updated_at TIMESTAMP
		)
	`)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Table 'books' created successfully")
}
