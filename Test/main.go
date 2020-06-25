package main

import (
	"github.com/gin-gonic/gin"
	"gqlgen-todos/Test/controllers"
	"gqlgen-todos/Test/models"
)

func main() {
	r := gin.Default()


	// Connect to database
	models.ConnectDatabase()

	// Routes
	r.GET("/books", controllers.FindBooks)
	r.GET("/books/:id", controllers.FindBook)
	r.POST("/books", controllers.CreateBook)
	r.PATCH("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)

	r.GET("/authors", controllers.FindAuthors)
	r.GET("/authors/:id", controllers.FindAuthor)
	r.POST("/authors", controllers.CreateAuthor)
	// Run the server
	r.Run(":4001")
}
