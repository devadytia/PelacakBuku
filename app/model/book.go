package book

type Book struct {
	Title  string `validate:"required"`
	Author string `validate:"required"`
	Rating int    `validate:"max=5,min=1"`
}
