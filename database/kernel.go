package main

import (
	database "PelacakBuku/database/migrations"
)

func main() {
	database.InitDatabase()

	database.CreateTable()
}
