// internal/router/router.go

package router

import (
	"github.com/gorilla/mux"
	"library-api/internal/handlers"
	"library-api/internal/repository"
	"library-api/internal/service"
)

// SetupRouter initializes the router
func SetupRouter() *mux.Router {
	// Initialize router
	r := mux.NewRouter()

	// Initialize dependencies
	repo := repository.NewInMemoryBookRepo()
	bookService := service.NewBookService(repo)
	bookHandler := handlers.NewBookHandler(bookService)

	// Register routes
	r.HandleFunc("/books", bookHandler.CreateBook).Methods("POST")
	r.HandleFunc("/books", bookHandler.GetAllBooks).Methods("GET")
	r.HandleFunc("/books/{id}", bookHandler.GetBookByID).Methods("GET")
	r.HandleFunc("/books/{id}", bookHandler.DeleteBook).Methods("DELETE")

	return r
}
