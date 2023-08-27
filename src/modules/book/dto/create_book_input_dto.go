package book

type CreateBookInput struct {
	Title  string `json:"title" validate:"required"`
	Author string `json:"author" validate:"required"`
}
