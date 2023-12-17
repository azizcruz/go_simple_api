package controllers

import (
	"bookstore/helpers"
	"bookstore/models"
	"bookstore/repositories"
	"bookstore/validators"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBooksRoute(c *gin.Context) {
	BookRepository := c.MustGet("bookRepository").(repositories.BookRepository)
	errorHandlers := &helpers.ErrorHandlers{}

	books, err := BookRepository.GetBooks()

	errorHandlers.HandleInternalError(c, err)

	c.JSON(http.StatusOK, books)
}

func GetBookByIdRoute(c *gin.Context) {
	BookRepository := c.MustGet("bookRepository").(repositories.BookRepository)
	errorHandlers := &helpers.ErrorHandlers{}

	id := c.Param("id")
	book, err := BookRepository.GetBookById(id)

	errorHandlers.HandleInternalError(c, err)

	if book.ID == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	fmt.Println(book)

	c.JSON(http.StatusOK, book)

}

func CreateBookRoute(c *gin.Context) {
	bookRepository := c.MustGet("bookRepository").(repositories.BookRepository)
	errorHandlers := &helpers.ErrorHandlers{}
	helpers := &helpers.Utils{}
	var requestData validators.CreateBookValidator

	if err := c.BindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	titleExists, err := bookRepository.CheckIfTitleExists(requestData.Title)

	if titleExists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title already exists"})
		return
	}

	book, err := bookRepository.CreateBook(models.Book{
		ID:     helpers.GenerateRandomID(),
		Author: requestData.Author,
		Genre:  requestData.Genre,
		Title:  requestData.Title,
	})

	errorHandlers.HandleInternalError(c, err)

	c.JSON(http.StatusOK, book)
}

func UpdateBookRoute(c *gin.Context) {
	bookRepository := c.MustGet("bookRepository").(repositories.BookRepository)
	errorHandlers := &helpers.ErrorHandlers{}
	var requestData validators.UpdateBookValidator

	if err := c.BindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")

	idExists, err := bookRepository.CheckIfIdExists(id)

	errorHandlers.HandleNotFoundError(c, idExists)

	updatedBook, err := bookRepository.UpdateBook(id, models.Book{
		Author: requestData.Author,
		Genre:  requestData.Genre,
		Title:  requestData.Title,
	})

	errorHandlers.HandleInternalError(c, err)

	c.JSON(http.StatusOK, updatedBook)
}

func DeleteBookRoute(c *gin.Context) {
	bookRepository := c.MustGet("bookRepository").(repositories.BookRepository)
	errorHandlers := &helpers.ErrorHandlers{}

	id := c.Param("id")

	idExists, err := bookRepository.CheckIfIdExists(id)

	errorHandlers.HandleNotFoundError(c, idExists)

	isDeleted, err := bookRepository.DeleteBook(id)

	errorHandlers.HandleInternalError(c, err)

	if isDeleted != false {
		c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
	}
}
