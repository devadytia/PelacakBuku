package main

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"gopkg.in/go-playground/validator.v9"
)

type Book struct {
	Title  string `validate:"required"`
	Author string `validate:"required"`
	Rating int    `validate:"gte=1,lte=5"`
}

func main() {
	validate := validator.New()
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		dsn := "root:@tcp(localhost:3306)/bookcollect"
		db, err := sql.Open("mysql", dsn)
		if err != nil {
			return err
		}
		defer db.Close()

		// Query untuk mendapatkan data dari tabel books
		rows, err := db.Query("SELECT title, author, rating FROM books")
		if err != nil {
			return err
		}
		defer rows.Close()

		// Loop melalui hasil query dan membangun slice Book
		var books []Book
		for rows.Next() {
			var book Book
			err := rows.Scan(&book.Title, &book.Author, &book.Rating)
			if err != nil {
				return err
			}
			books = append(books, book)
		}

		return c.Render("form", fiber.Map{"Books": books})
	})

	app.Post("/submit", func(c *fiber.Ctx) error {
		title := c.FormValue("title")
		author := c.FormValue("author")
		ratingStr := c.FormValue("rating")

		rating, err := strconv.Atoi(ratingStr)
		if err != nil {
			return err
		}

		book := Book{
			Title:  title,
			Author: author,
			Rating: rating,
		}

		errVal := validate.Struct(book)
		if errVal != nil {
			return errVal
		}

		fmt.Println("Data is valid!")

		err = saveBookToDatabase(book)
		if err != nil {
			return err
		}

		return c.Redirect("/")
	})

	app.Listen(":9998")
}

func saveBookToDatabase(book Book) error {
	dsn := "root:@tcp(localhost:3306)/bookcollect"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	defer db.Close()

	currentTime := time.Now()

	// Simpan data buku ke dalam tabel
	_, err = db.Exec("INSERT INTO books (title, author, rating, created_at, updated_at) VALUES (?, ?, ?, ?, ?)", book.Title, book.Author, book.Rating, currentTime, currentTime)
	if err != nil {
		return err
	}

	return nil
}
