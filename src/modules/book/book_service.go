package book

import (
	dto "book-api/src/modules/book/dto"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type BookService interface {
	GetAll() []Book
	GetByID(id int64) Book
	Create(ctx *gin.Context) (*Book, error)
	Update(ctx *gin.Context) (*Book, error)
	Delete(ctx *gin.Context) (*Book, error)
}

type BookServiceImpl struct {
	bookRepository BookRepository
}

func NewBookService(bookRepository BookRepository) BookService {
	return &BookServiceImpl{bookRepository}
}

func (bs *BookServiceImpl) GetAll() []Book {
	return bs.bookRepository.FindAll()
}

func (bs *BookServiceImpl) GetByID(id int64) Book {
	return bs.bookRepository.FindByID(int64(id))
}

func (bs *BookServiceImpl) Create(ctx *gin.Context) (*Book, error) {
	var input dto.CreateBookInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		return nil, err
	}

	validate := validator.New()

	err := validate.Struct(input)

	if err != nil {
		return nil, err
	}

	book := Book{
		Title:  input.Title,
		Author: input.Author,
	}

	bookResult, err := bs.bookRepository.Save(book)

	if err != nil {
		return nil, err
	}
	return bookResult, nil
}

func (bs *BookServiceImpl) Update(ctx *gin.Context) (*Book, error) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	var input dto.UpdateBookInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		return nil, err
	}

	validate := validator.New()

	err := validate.Struct(input)

	if err != nil {
		return nil, err
	}

	book := Book{
		ID:     int64(id),
		Title:  input.Title,
		Author: input.Author,
	}

	bookResult, err := bs.bookRepository.Update(book)

	if err != nil {
		return nil, err
	}
	return bookResult, nil

}

func (bs *BookServiceImpl) Delete(ctx *gin.Context) (*Book, error) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	book := Book{
		ID: int64(id),
	}

	bookResult, err := bs.bookRepository.Delete(book)

	if err != nil {
		return nil, err
	}
	return bookResult, nil
}
