package main

import (
	helper "PelacakBuku/app/helper"
	route "PelacakBuku/route"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	engine := html.New("resources/views", ".html")

	helper.RegisterFuncs(engine)

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	route.WebRoutes(app, engine)

	app.Listen(":9998")
}
