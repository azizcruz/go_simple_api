package middlewares

import (
	"bookstore/repositories"

	"github.com/gin-gonic/gin"
)

func InjectBookRepositoryMiddleware(bookRepository repositories.BookRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("bookRepository", bookRepository)
		c.Next()
	}
}
