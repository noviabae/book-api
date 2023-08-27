package main

import (
	"book-api/src/config"
	"book-api/src/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	db := config.DB()

	routes.API(router, db)
	router.Run()
}
