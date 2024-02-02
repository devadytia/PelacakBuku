package helper

import (
	"github.com/gofiber/template/html/v2"
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
