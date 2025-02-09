package tests

import (
	"library-api/internal/models"
	"library-api/internal/repository"
	"testing"
)

func TestCreateBook(t *testing.T) {
	repo := repository.NewInMemoryBookRepo()

	book := &models.Book{
		Title:         "Test Book",
		Author:        "Test Author",
		PublishedYear: 2022,
	}

	err := repo.Create(book)
	if err != nil {
		t.Errorf("Failed to create book: %v", err)
	}

	if book.ID == 0 {
		t.Error("Expected book ID to be set after creation")
	}
}

func TestGetBookByID(t *testing.T) {
	repo := repository.NewInMemoryBookRepo()

	book := &models.Book{
		Title:         "Test Book",
		Author:        "Test Author",
		PublishedYear: 2022,
	}

	// Create a book first
	err := repo.Create(book)
	if err != nil {
		t.Fatalf("Failed to create book: %v", err)
	}

	// Try to retrieve the book
	retrieved, err := repo.GetByID(book.ID)
	if err != nil {
		t.Errorf("Failed to get book by ID: %v", err)
	}

	if retrieved.Title != book.Title {
		t.Errorf("Expected book title %s, got %s", book.Title, retrieved.Title)
	}
}
func TestDeleteBook(t *testing.T) {
	repo := repository.NewInMemoryBookRepo()

	book := &models.Book{
		Title:         "Delete Test Book",
		Author:        "Delete Test Author",
		PublishedYear: 2022,
	}

	// Create a book first
	err := repo.Create(book)
	if err != nil {
		t.Fatalf("Failed to create book: %v", err)
	}

	// Try to delete the book
	err = repo.Delete(book.ID)
	if err != nil {
		t.Errorf("Failed to delete book: %v", err)
	}

	// Verify the book is deleted
	_, err = repo.GetByID(book.ID)
	if err == nil {
		t.Error("Expected error when getting deleted book")
	}
}

func TestGetAllBooks(t *testing.T) {
	repo := repository.NewInMemoryBookRepo()

	// Create multiple books
	books := []*models.Book{
		{
			Title:         "Book 1",
			Author:        "Author 1",
			PublishedYear: 2022,
		},
		{
			Title:         "Book 2",
			Author:        "Author 2",
			PublishedYear: 2023,
		},
	}

	for _, book := range books {
		err := repo.Create(book)
		if err != nil {
			t.Fatalf("Failed to create book: %v", err)
		}
	}

	// Try to retrieve all books
	retrievedBooks, err := repo.GetAll()
	if err != nil {
		t.Errorf("Failed to get all books: %v", err)
	}

	if len(retrievedBooks) != len(books) {
		t.Errorf("Expected %d books, got %d", len(books), len(retrievedBooks))
	}
}

func TestSearchByAuthor(t *testing.T) {
	repo := repository.NewInMemoryBookRepo()

	// Create books with the same author
	author := "Test Author"
	books := []*models.Book{
		{
			Title:         "Book 1",
			Author:        author,
			PublishedYear: 2022,
		},
		{
			Title:         "Book 2",
			Author:        author,
			PublishedYear: 2023,
		},
		{
			Title:         "Book 3",
			Author:        "Different Author",
			PublishedYear: 2023,
		},
	}

	for _, book := range books {
		err := repo.Create(book)
		if err != nil {
			t.Fatalf("Failed to create book: %v", err)
		}
	}

	// Search for books by author
	foundBooks, err := repo.SearchByAuthor(author)
	if err != nil {
		t.Errorf("Failed to search books by author: %v", err)
	}

	if len(foundBooks) != 2 {
		t.Errorf("Expected 2 books by author %s, got %d", author, len(foundBooks))
	}
}

func TestSearchNonExistentAuthor(t *testing.T) {
	repo := repository.NewInMemoryBookRepo()

	// Search for books by non-existent author
	_, err := repo.SearchByAuthor("Non Existent Author")
	if err == nil {
		t.Error("Expected error when searching for non-existent author")
	}
}

func TestDeleteNonExistentBook(t *testing.T) {
	repo := repository.NewInMemoryBookRepo()

	// Try to delete non-existent book
	err := repo.Delete(999)
	if err == nil {
		t.Error("Expected error when deleting non-existent book")
	}
}

func TestGetNonExistentBook(t *testing.T) {
	repo := repository.NewInMemoryBookRepo()

	// Try to get non-existent book
	_, err := repo.GetByID(999)
	if err == nil {
		t.Error("Expected error when getting non-existent book")
	}
}

func TestCreateDuplicateBook(t *testing.T) {
	repo := repository.NewInMemoryBookRepo()

	book := &models.Book{
		Title:         "Test Book",
		Author:        "Test Author",
		PublishedYear: 2022,
	}

	// Create first book
	err := repo.Create(book)
	if err != nil {
		t.Fatalf("Failed to create first book: %v", err)
	}

	// Try to create duplicate book
	duplicateBook := &models.Book{
		Title:         "Test Book",
		Author:        "Test Author",
		PublishedYear: 2022,
	}

	err = repo.Create(duplicateBook)
	if err != nil {
		t.Errorf("Failed to create duplicate book: %v", err)
	}

	if book.ID == duplicateBook.ID {
		t.Error("Duplicate book should have different ID")
	}
}
