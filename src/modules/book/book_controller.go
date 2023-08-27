package book

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookController struct {
	bookService BookService
	ctx         *gin.Context
}

func NewBookController(bookService BookService, ctx *gin.Context) BookController {
	return BookController{bookService, ctx}
}

func (bc *BookController) Index(ctx *gin.Context) {
	data := bc.bookService.GetAll()

	ctx.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"data":   data,
	})
}

func (bc *BookController) GetByID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	data := bc.bookService.GetByID(int64(id))

	ctx.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"data":   data,
	})
}

func (bc *BookController) Create(ctx *gin.Context) {
	data, err := bc.bookService.Create(ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"data":   err,
		})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"data":   data,
	})
}

func (bc *BookController) Update(ctx *gin.Context) {
	data, err := bc.bookService.Update(ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"data":   err,
		})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"data":   data,
	})
}

func (bc *BookController) Delete(ctx *gin.Context) {
	data, err := bc.bookService.Delete(ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"data":   err,
		})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"data":   data,
	})
}
