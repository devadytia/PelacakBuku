package main

import (
	"PelacakBuku/internal/database"
)

func main() {
	database.InitDatabase()

	database.CreateDatabase()

	database.UseDatabase()

	database.CreateTable()
}
