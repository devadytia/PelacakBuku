package bookservice

import (
	book "PelacakBuku/app/model"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

func StoreBook(book book.Book) error {
	dsn := "user=postgres dbname=bookcollect sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return err
	}
	defer db.Close()

	currentTime := time.Now()

	_, err = db.Exec("INSERT INTO books (title, author, rating, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)",
		book.Title, book.Author, book.Rating, currentTime, currentTime)
	if err != nil {
		return err
	}

	fmt.Println("Book stored successfully")
	return nil
}
