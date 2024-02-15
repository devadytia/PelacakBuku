package bookservice

import (
	helper "PelacakBuku/app/helper"
	book "PelacakBuku/app/model"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

func GetBook() ([]book.Book, error) {
	db, err := helper.Connect()

	if err != nil {
		return nil, err
	}

	rows, err := db.Query("SELECT title, author, rating FROM books ORDER BY id DESC")

	if err != nil {
		return nil, err
	}

	var books []book.Book

	for rows.Next() {
		var book book.Book
		err := rows.Scan(&book.Title, &book.Author, &book.Rating)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	return books, err
}

func StoreBook(book book.Book) error {
	db, err := helper.Connect()

	if err != nil {
		return err
	}

	currentTime := time.Now()

	_, err = db.Exec("INSERT INTO books (title, author, rating, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)",
		book.Title, book.Author, book.Rating, currentTime, currentTime)
	if err != nil {
		return err
	}

	fmt.Println("Book stored successfully")
	return nil
}
