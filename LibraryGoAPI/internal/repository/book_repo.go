package repository

import (
	"errors"
	"library-api/internal/interfaces"
	"library-api/internal/models"
	"library-api/pkg/logger"
	"sync"
	"sync/atomic"
)

var _ interfaces.BookRepository = (*InMemoryBookRepo)(nil)

type InMemoryBookRepo struct {
	mutex  sync.RWMutex
	books  map[int]*models.Book
	lastID int64
	logger *logger.Logger
}

func NewInMemoryBookRepo() *InMemoryBookRepo {
	return &InMemoryBookRepo{
		books:  make(map[int]*models.Book),
		lastID: 0,
		logger: logger.NewLogger(),
	}
}

func (r *InMemoryBookRepo) Create(book *models.Book) error {
	r.logger.Info("Attempting to create new book: Title=%s, Author=%s", book.Title, book.Author)

	r.mutex.Lock()
	defer r.mutex.Unlock()
	r.logger.Debug("Creating new book: %+v", book)

	id := int(atomic.AddInt64(&r.lastID, 1))
	book.ID = id
	r.books[book.ID] = book
	r.logger.Info("Successfully created book with ID: %d", book.ID)

	return nil
}

func (r *InMemoryBookRepo) GetAll() ([]*models.Book, error) {
	r.logger.Info("Retrieving all books")

	r.mutex.RLock()
	defer r.mutex.RUnlock()

	books := make([]*models.Book, 0, len(r.books))
	for _, book := range r.books {
		books = append(books, book)
	}
	r.logger.Info("Successfully retrieved %d books", len(books))
	r.logger.Debug("Retrieved books: %+v", books)
	return books, nil
}

func (r *InMemoryBookRepo) GetByID(id int) (*models.Book, error) {
	r.logger.Info("Attempting to retrieve book with ID: %d", id)

	r.mutex.RLock()
	defer r.mutex.RUnlock()

	book, exists := r.books[id]
	if !exists {
		r.logger.Error("Book not found with ID: %d", id)

		return nil, errors.New("book not found")
	}
	r.logger.Debug("Retrieved book details: %+v", book)
	r.logger.Info("Successfully retrieved book with ID: %d", id)
	return book, nil
}

func (r *InMemoryBookRepo) Delete(id int) error {
	r.logger.Info("Attempting to delete book with ID: %d", id)

	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.books[id]; !exists {
		r.logger.Error("Failed to delete - book not found with ID: %d", id)

		return errors.New("book not found")
	}
	delete(r.books, id)
	r.logger.Info("Successfully deleted book with ID: %d", id)

	return nil
}

func (r *InMemoryBookRepo) SearchByAuthor(author string) ([]*models.Book, error) {
	r.logger.Info("Searching for books by author: %s", author)

	r.mutex.RLock()
	defer r.mutex.RUnlock()

	var result []*models.Book
	for _, book := range r.books {
		if book.Author == author {
			result = append(result, book)
		}
	}
	if len(result) == 0 {
		r.logger.Error("No books found for author: %s", author)

		return nil, errors.New("author not found")
	}
	r.logger.Debug("Found %d books by author: %s", len(result), author)
	r.logger.Info("Successfully retrieved books by author: %s", author)
	return result, nil
}
