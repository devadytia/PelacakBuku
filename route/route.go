package route

import (
	bookcontroller "PelacakBuku/app/http/controller"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func WebRoutes(app *fiber.App, engine *html.Engine) {
	app.Static("/css", "./public/css")
	app.Static("/img", "./public/img")

	app.Get("/", bookcontroller.Index)
	app.Post("/submit", bookcontroller.Store)
}
