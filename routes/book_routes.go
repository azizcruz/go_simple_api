package routes

import (
	"bookstore/controllers"
	"bookstore/middlewares"
	"bookstore/repositories"

	"github.com/gin-gonic/gin"
)

func RegisterBookRoutes(r *gin.Engine, bookRepository repositories.BookRepository) {
	r.Use(middlewares.InjectBookRepositoryMiddleware(bookRepository))

	r.GET("/books/:id", controllers.GetBookByIdRoute)
	r.PUT("/books/:id", controllers.UpdateBookRoute)

	r.GET("/books", controllers.GetBooksRoute)
	r.POST("/books", controllers.CreateBookRoute)

	r.DELETE("/books/:id", controllers.DeleteBookRoute)
}
