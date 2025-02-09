package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"library-api/internal/interfaces"
	"library-api/internal/models"
	"library-api/pkg/api"
	"net/http"
	"strconv"
)

type BookHandler struct {
	service interfaces.BookService
}

func NewBookHandler(service interfaces.BookService) interfaces.BookHandler {
	return &BookHandler{service: service}
}

func handlePanic(w http.ResponseWriter) {
	if r := recover(); r != nil {
		writeJSON(w, http.StatusInternalServerError, api.Response{
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf("Internal server error: %v", r),
		})
	}
}

func (h *BookHandler) CreateBook(w http.ResponseWriter, r *http.Request) {
	defer handlePanic(w)

	if r.Body == nil {
		writeJSON(w, http.StatusBadRequest, api.Response{
			Status:  http.StatusBadRequest,
			Message: "Request body is empty",
		})
		return
	}

	var book models.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		writeJSON(w, http.StatusBadRequest, api.Response{
			Status:  http.StatusBadRequest,
			Message: fmt.Sprintf("Invalid request body: %v", err),
		})
		return
	}

	// Validate book data
	if err := validateBook(&book); err != nil {
		writeJSON(w, http.StatusBadRequest, api.Response{
			Status:  http.StatusBadRequest,
			Message: fmt.Sprintf("Invalid book data: %v", err),
		})
		return
	}

	if err := h.service.CreateBook(&book); err != nil {
		writeJSON(w, http.StatusInternalServerError, api.Response{
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf("Failed to create book: %v", err),
		})
		return
	}

	writeJSON(w, http.StatusCreated, api.Response{
		Status:  http.StatusCreated,
		Message: "Book created successfully",
		Data:    book,
	})
}

func (h *BookHandler) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	defer handlePanic(w)

	author := r.URL.Query().Get("author")
	var books []*models.Book
	var err error

	if author != "" {
		books, err = h.service.SearchBooksByAuthor(author)
		if err != nil {
			status := http.StatusInternalServerError
			message := "Failed to retrieve books"

			if err.Error() == "author not found" {
				status = http.StatusNotFound
				message = "No books found for the specified author"
			}

			writeJSON(w, status, api.Response{
				Status:  status,
				Message: fmt.Sprintf("%s: %v", message, err),
			})
			return
		}
	} else {
		books, err = h.service.GetAllBooks()
		if err != nil {
			writeJSON(w, http.StatusInternalServerError, api.Response{
				Status:  http.StatusInternalServerError,
				Message: fmt.Sprintf("Failed to retrieve books: %v", err),
			})
			return
		}
	}

	if len(books) == 0 {
		writeJSON(w, http.StatusNotFound, api.Response{
			Status:  http.StatusNotFound,
			Message: "No books found",
		})
		return
	}

	writeJSON(w, http.StatusOK, api.Response{
		Status:  http.StatusOK,
		Message: "Books retrieved successfully",
		Data:    books,
	})
}

func (h *BookHandler) GetBookByID(w http.ResponseWriter, r *http.Request) {
	defer handlePanic(w)

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		writeJSON(w, http.StatusBadRequest, api.Response{
			Status:  http.StatusBadRequest,
			Message: fmt.Sprintf("Invalid book ID: %v", err),
		})
		return
	}

	if id <= 0 {
		writeJSON(w, http.StatusBadRequest, api.Response{
			Status:  http.StatusBadRequest,
			Message: "Book ID must be positive",
		})
		return
	}

	book, err := h.service.GetBookByID(id)
	if err != nil {
		writeJSON(w, http.StatusNotFound, api.Response{
			Status:  http.StatusNotFound,
			Message: fmt.Sprintf("Book not found: %v", err),
		})
		return
	}

	writeJSON(w, http.StatusOK, api.Response{
		Status:  http.StatusOK,
		Message: "Book retrieved successfully",
		Data:    book,
	})
}

func (h *BookHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	defer handlePanic(w)

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		writeJSON(w, http.StatusBadRequest, api.Response{
			Status:  http.StatusBadRequest,
			Message: fmt.Sprintf("Invalid book ID: %v", err),
		})
		return
	}

	if id <= 0 {
		writeJSON(w, http.StatusBadRequest, api.Response{
			Status:  http.StatusBadRequest,
			Message: "Book ID must be positive",
		})
		return
	}

	if err := h.service.DeleteBook(id); err != nil {
		writeJSON(w, http.StatusNotFound, api.Response{
			Status:  http.StatusNotFound,
			Message: fmt.Sprintf("Failed to delete book: %v", err),
		})
		return
	}

	writeJSON(w, http.StatusOK, api.Response{
		Status:  http.StatusOK,
		Message: "Book deleted successfully",
	})
}

func validateBook(book *models.Book) error {
	if book == nil {
		return fmt.Errorf("book cannot be nil")
	}
	if book.Title == "" {
		return fmt.Errorf("book title is required")
	}
	if book.Author == "" {
		return fmt.Errorf("book author is required")
	}
	return nil
}

func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Failed to encode JSON response", http.StatusInternalServerError)
	}
}
