package bookcontroller

import (
	book "PelacakBuku/app/model"
	bookservice "PelacakBuku/app/services"

	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"gopkg.in/go-playground/validator.v9"
)

var validate = validator.New()

func Index(c *fiber.Ctx) error {
	books, err := bookservice.GetBook()

	if err != nil {
		return err
	}

	return c.Render("book", fiber.Map{"Books": books})
}

func Store(c *fiber.Ctx) error {
	title := c.FormValue("title")
	author := c.FormValue("author")
	ratingStr := c.FormValue("rating")

	rating, err := strconv.Atoi(ratingStr)
	if err != nil {
		return err
	}

	book := book.Book{
		Title:  title,
		Author: author,
		Rating: rating,
	}

	errVal := validate.Struct(book)
	if errVal != nil {
		return errVal
	}

	fmt.Println("Data is valid!")

	err = bookservice.StoreBook(book)
	if err != nil {
		return err
	}

	return c.Redirect("/")
}
