package service

import (
	"library-api/internal/interfaces"
	"library-api/internal/models"
)

type BookService struct {
	repo interfaces.BookRepository
}

func NewBookService(repo interfaces.BookRepository) interfaces.BookService {
	return &BookService{repo: repo}
}

func (s *BookService) CreateBook(book *models.Book) error {
	return s.repo.Create(book)
}

func (s *BookService) GetAllBooks() ([]*models.Book, error) {
	return s.repo.GetAll()
}

func (s *BookService) GetBookByID(id int) (*models.Book, error) {
	return s.repo.GetByID(id)
}

func (s *BookService) DeleteBook(id int) error {
	return s.repo.Delete(id)
}

func (s *BookService) SearchBooksByAuthor(author string) ([]*models.Book, error) {
	return s.repo.SearchByAuthor(author)
}
