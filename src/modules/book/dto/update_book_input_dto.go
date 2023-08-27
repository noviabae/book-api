package book

type UpdateBookInput struct {
	Title  string `json:"title" validate:"required"`
	Author string `json:"author" validate:"required"`
}
