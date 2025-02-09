package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"library-api/internal/interfaces"
	"library-api/internal/models"
	"library-api/pkg/api"
	"library-api/pkg/logger"
	"net/http"
	"strconv"
)

type BookHandler struct {
	service interfaces.BookService
	logger  *logger.Logger
}

func NewBookHandler(service interfaces.BookService) interfaces.BookHandler {
	return &BookHandler{
		service: service,
		logger:  logger.NewLogger(),
	}
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
	h.logger.Info("Handling CreateBook request")

	defer handlePanic(w)

	if r.Body == nil {
		h.logger.Error("Request body is empty")

		writeJSON(w, http.StatusBadRequest, api.Response{
			Status:  http.StatusBadRequest,
			Message: "Request body is empty",
		})
		return
	}

	var book models.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		h.logger.Error("Failed to decode request body: %v", err)

		writeJSON(w, http.StatusBadRequest, api.Response{
			Status:  http.StatusBadRequest,
			Message: fmt.Sprintf("Invalid request body: %v", err),
		})
		return
	}
	h.logger.Debug("Received book data: %+v", book)

	// Validate book data
	if err := validateBook(&book); err != nil {
		h.logger.Error("Book validation failed: %v", err)

		writeJSON(w, http.StatusBadRequest, api.Response{
			Status:  http.StatusBadRequest,
			Message: fmt.Sprintf("Invalid book data: %v", err),
		})
		return
	}

	if err := h.service.CreateBook(&book); err != nil {
		h.logger.Error("Failed to create book: %v", err)

		writeJSON(w, http.StatusInternalServerError, api.Response{
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf("Failed to create book: %v", err),
		})
		return
	}
	h.logger.Info("Successfully created book with ID: %d", book.ID)

	writeJSON(w, http.StatusCreated, api.Response{
		Status:  http.StatusCreated,
		Message: "Book created successfully",
		Data:    book,
	})
}

func (h *BookHandler) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("Handling GetAllBooks request")

	defer handlePanic(w)

	author := r.URL.Query().Get("author")
	var books []*models.Book
	var err error

	if author != "" {
		h.logger.Debug("Searching books by author: %s", author)

		books, err = h.service.SearchBooksByAuthor(author)
		if err != nil {
			status := http.StatusInternalServerError
			message := "Failed to retrieve books"

			if err.Error() == "author not found" {
				h.logger.Info("No books found for author: %s", author)

				status = http.StatusNotFound
				message = "No books found for the specified author"
			} else {
				h.logger.Error("Error searching books by author: %v", err)
			}

			writeJSON(w, status, api.Response{
				Status:  status,
				Message: fmt.Sprintf("%s: %v", message, err),
			})
			return
		}
	} else {
		h.logger.Debug("Retrieving all books")

		books, err = h.service.GetAllBooks()
		if err != nil {
			h.logger.Error("Failed to retrieve all books: %v", err)

			writeJSON(w, http.StatusInternalServerError, api.Response{
				Status:  http.StatusInternalServerError,
				Message: fmt.Sprintf("Failed to retrieve books: %v", err),
			})
			return
		}
	}

	if len(books) == 0 {
		h.logger.Info("No books found")

		writeJSON(w, http.StatusNotFound, api.Response{
			Status:  http.StatusNotFound,
			Message: "No books found",
		})
		return
	}
	h.logger.Info("Successfully retrieved %d books", len(books))

	writeJSON(w, http.StatusOK, api.Response{
		Status:  http.StatusOK,
		Message: "Books retrieved successfully",
		Data:    books,
	})
}

func (h *BookHandler) GetBookByID(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("Handling GetBookByID request")

	defer handlePanic(w)

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		h.logger.Error("Invalid book ID format: %s", vars["id"])

		writeJSON(w, http.StatusBadRequest, api.Response{
			Status:  http.StatusBadRequest,
			Message: fmt.Sprintf("Invalid book ID: %v", err),
		})
		return
	}

	if id <= 0 {
		h.logger.Error("Invalid book ID: %d", id)

		writeJSON(w, http.StatusBadRequest, api.Response{
			Status:  http.StatusBadRequest,
			Message: "Book ID must be positive",
		})
		return
	}
	h.logger.Debug("Fetching book with ID: %d", id)

	book, err := h.service.GetBookByID(id)
	if err != nil {
		writeJSON(w, http.StatusNotFound, api.Response{
			Status:  http.StatusNotFound,
			Message: fmt.Sprintf("Book not found: %v", err),
		})
		return
	}
	h.logger.Info("Successfully retrieved book with ID: %d", id)

	writeJSON(w, http.StatusOK, api.Response{
		Status:  http.StatusOK,
		Message: "Book retrieved successfully",
		Data:    book,
	})
}

func (h *BookHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("Handling DeleteBook request")

	defer handlePanic(w)

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		h.logger.Error("Invalid book ID format: %s", vars["id"])

		writeJSON(w, http.StatusBadRequest, api.Response{
			Status:  http.StatusBadRequest,
			Message: fmt.Sprintf("Invalid book ID: %v", err),
		})
		return
	}

	if id <= 0 {
		h.logger.Error("Invalid book ID: %d", id)

		writeJSON(w, http.StatusBadRequest, api.Response{
			Status:  http.StatusBadRequest,
			Message: "Book ID must be positive",
		})
		return
	}
	h.logger.Debug("Attempting to delete book with ID: %d", id)

	if err := h.service.DeleteBook(id); err != nil {
		h.logger.Error("Failed to delete book with ID %d: %v", id, err)

		writeJSON(w, http.StatusNotFound, api.Response{
			Status:  http.StatusNotFound,
			Message: fmt.Sprintf("Failed to delete book: %v", err),
		})
		return
	}
	h.logger.Info("Successfully deleted book with ID: %d", id)

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
