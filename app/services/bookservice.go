package bookservice

import (
	book "PelacakBuku/app/model"
	"database/sql"
	"time"
)

func StoreBook(book book.Book) error {
	dsn := "root:@tcp(localhost:3306)/bookcollect"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	defer db.Close()

	currentTime := time.Now()

	_, err = db.Exec("INSERT INTO books (title, author, rating, created_at, updated_at) VALUES (?, ?, ?, ?, ?)", book.Title, book.Author, book.Rating, currentTime, currentTime)
	if err != nil {
		return err
	}

	return nil
}
