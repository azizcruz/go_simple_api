package repositories

import (
	"bookstore/models"
)

type BookRepository interface {
	GetBooks() ([]models.Book, error)
	GetBookById(id string) (models.Book, error)
	CreateBook(book models.Book) (models.Book, error)
	UpdateBook(id string, book models.Book) (models.Book, error)
	DeleteBook(id string) (bool, error)
	CheckIfTitleExists(title string) (bool, error)
	CheckIfIdExists(id string) (bool, error)
}

type InMemoryBookRepository struct {
	books []models.Book
}

func NewInMemoryBookRepository() *InMemoryBookRepository {
	return &InMemoryBookRepository{
		books: []models.Book{},
	}
}

func (i *InMemoryBookRepository) CreateBook(book models.Book) (models.Book, error) {
	i.books = append(i.books, book)

	return book, nil
}

func (i *InMemoryBookRepository) GetBooks() ([]models.Book, error) {
	return i.books, nil
}

func (i *InMemoryBookRepository) GetBookById(id string) (models.Book, error) {
	for _, book := range i.books {
		if book.ID == id {
			return book, nil
		}
	}

	return models.Book{}, nil
}

func (i *InMemoryBookRepository) UpdateBook(id string, book models.Book) (models.Book, error) {
	for index, b := range i.books {
		if b.ID == id {
			if book.Author != "" {
				b.Author = book.Author
			}
			if book.Genre != "" {
				b.Genre = book.Genre
			}
			if book.Title != "" {
				b.Title = book.Title
			}

			i.books[index] = b

			return b, nil
		}
	}
	return models.Book{}, nil
}

func (i *InMemoryBookRepository) DeleteBook(id string) (bool, error) {
	for index, b := range i.books {
		if b.ID == id {
			newBooks := append(i.books[:index], i.books[index+1:]...)
			i.books = newBooks
			return true, nil
		}
	}

	return false, nil
}

func (i *InMemoryBookRepository) CheckIfTitleExists(title string) (bool, error) {
	for _, book := range i.books {
		if book.Title == title {
			return true, nil
		}
	}

	return false, nil
}

func (i *InMemoryBookRepository) CheckIfIdExists(id string) (bool, error) {
	for _, book := range i.books {
		if book.ID == id {
			return true, nil
		}
	}

	return false, nil
}
