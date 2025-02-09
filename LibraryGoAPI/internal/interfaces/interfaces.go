// internal/interfaces/interfaces.go

package interfaces

import (
	"library-api/internal/models"
	"net/http"
)

type BookRepository interface {
	Create(book *models.Book) error
	GetAll() ([]*models.Book, error)
	GetByID(id int) (*models.Book, error)
	Delete(id int) error
	SearchByAuthor(author string) ([]*models.Book, error)
}

type BookService interface {
	CreateBook(book *models.Book) error
	GetAllBooks() ([]*models.Book, error)
	GetBookByID(id int) (*models.Book, error)
	DeleteBook(id int) error
	SearchBooksByAuthor(author string) ([]*models.Book, error)
}

type BookHandler interface {
	CreateBook(w http.ResponseWriter, r *http.Request)
	GetAllBooks(w http.ResponseWriter, r *http.Request)
	GetBookByID(w http.ResponseWriter, r *http.Request)
	DeleteBook(w http.ResponseWriter, r *http.Request)
}
