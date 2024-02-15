package helper

import (
	"database/sql"
	"sync"

	"github.com/gofiber/template/html/v2"
)

var (
	db     *sql.DB
	dbOnce sync.Once
)

func Until(n int) []int {
	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = i
	}
	return result
}

func RegisterFuncs(engine *html.Engine) {
	engine.AddFunc("Until", Until)
}

func Connect() (*sql.DB, error) {
	var err error
	dbOnce.Do(func() {
		dsn := "user=postgres dbname=bookcollect sslmode=disable"
		db, err = sql.Open("postgres", dsn)
	})

	return db, err
}
