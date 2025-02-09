package service

import (
	"library-api/internal/interfaces"
	"library-api/internal/models"
	"library-api/pkg/logger"
)

type BookService struct {
	repo   interfaces.BookRepository
	logger *logger.Logger
}

func NewBookService(repo interfaces.BookRepository) interfaces.BookService {
	logger := logger.NewLogger()
	logger.Info("Initializing new BookService")
	return &BookService{
		repo:   repo,
		logger: logger,
	}
}

func (s *BookService) CreateBook(book *models.Book) error {
	s.logger.Info("Service: Creating new book - Title: %s, Author: %s", book.Title, book.Author)
	return s.repo.Create(book)
}

func (s *BookService) GetAllBooks() ([]*models.Book, error) {
	s.logger.Info("Service: Retrieving all books")
	return s.repo.GetAll()
}

func (s *BookService) GetBookByID(id int) (*models.Book, error) {
	s.logger.Info("Service: Retrieving book with ID: %d", id)
	return s.repo.GetByID(id)
}

func (s *BookService) DeleteBook(id int) error {
	s.logger.Info("Service: Attempting to delete book with ID: %d", id)
	return s.repo.Delete(id)
}

func (s *BookService) SearchBooksByAuthor(author string) ([]*models.Book, error) {
	s.logger.Info("Service: Searching for books by author: %s", author)

	return s.repo.SearchByAuthor(author)
}
