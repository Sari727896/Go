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
