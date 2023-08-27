package routes

import (
	book "book-api/src/modules/book"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	ctx *gin.Context
)

func API(router *gin.Engine, db *gorm.DB) {
	bookRepository := book.NewBookRepositoryImplement(db)
	bookService := book.NewBookService(bookRepository)
	bookController := book.NewBookController(bookService, ctx)

	v1 := router.Group("/api/v1")
	{
		v1.GET("/books", bookController.Index)
		v1.GET("/books/:id", bookController.GetByID)
		v1.POST("/books", bookController.Create)
		v1.PATCH("/books/:id", bookController.Update)
		v1.DELETE("/books/:id", bookController.Delete)
	}
}
