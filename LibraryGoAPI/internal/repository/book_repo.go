package repository

import (
	"errors"
	"library-api/internal/interfaces"
	"library-api/internal/models"
	"sync"
	"sync/atomic"
)

var _ interfaces.BookRepository = (*InMemoryBookRepo)(nil)

type InMemoryBookRepo struct {
	mutex  sync.RWMutex
	books  map[int]*models.Book
    lastID  int64
}

func NewInMemoryBookRepo() *InMemoryBookRepo {
	return &InMemoryBookRepo{
		books:  make(map[int]*models.Book),
        lastID: 0,
	}
}

func (r *InMemoryBookRepo) Create(book *models.Book) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	id := int(atomic.AddInt64(&r.lastID, 1))
    book.ID = id
    r.books[book.ID] = book
    return nil
}

func (r *InMemoryBookRepo) GetAll() ([]*models.Book, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	books := make([]*models.Book, 0, len(r.books))
	for _, book := range r.books {
		books = append(books, book)
	}
	return books, nil
}

func (r *InMemoryBookRepo) GetByID(id int) (*models.Book, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	book, exists := r.books[id]
	if !exists {
		return nil, errors.New("book not found")
	}
	return book, nil
}

func (r *InMemoryBookRepo) Delete(id int) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.books[id]; !exists {
		return errors.New("book not found")
	}
	delete(r.books, id)
	return nil
}

func (r *InMemoryBookRepo) SearchByAuthor(author string) ([]*models.Book, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	var result []*models.Book
	for _, book := range r.books {
		if book.Author == author {
			result = append(result, book)
		}
	}
	if len(result) == 0 {
		return nil, errors.New("author not found")
	}
	return result, nil
}
