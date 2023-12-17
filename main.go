package main

import (
	"bookstore/repositories"
	"bookstore/routes"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	BookRepository := repositories.NewInMemoryBookRepository()

	routes.RegisterBookRoutes(r, BookRepository)

	fmt.Println("Listening on port 8080")

	r.Run()
}
