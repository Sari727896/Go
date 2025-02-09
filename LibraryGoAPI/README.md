# Library Management API

A REST API implementation for a library system that manages books, built with Go as per the coding assessment requirements.

## Requirements Implementation

### 1. REST API Endpoints
All required endpoints have been implemented:

- **POST /books** - Add a new book
  ```json
  {
    "title": "Book Title",
    "author": "Author Name",
    "publishedYear": 2022
  }
  ```
- **GET /books** - Retrieve all books
- **GET /books/{id}** - Retrieve a book by ID
- **DELETE /books/{id}** - Delete a book by ID

### 2. Router Implementation
- Using Gorilla Mux for routing
- Routes are configured in `internal/router/router.go`

### 3. Data Storage
- Implemented in-memory storage using a thread-safe map
- Located in `internal/repository/book_repo.go`

### 4. Error Handling
Implemented comprehensive error handling:
- 404 Not Found for non-existent books
- 400 Bad Request for invalid inputs
- 500 Internal Server Error for server-side issues
- Appropriate status codes for all operations

### Bonus Features Implemented

1. **Search Functionality**
   - Added GET /books?author=Author Name endpoint
   - Implements author-based search

2. **Project Structure**
   ```
   library-api/
   ├── internal/
   │   ├── handlers/      # HTTP handlers
   │   ├── interfaces/    # Interfaces
   │   ├── models/        # Data models
   │   ├── repository/    # Data storage
   │   ├── router/        # Router setup
   │   └── service/       # Business logic
   └── pkg/
       └── api/          # Response structures
   ```

## Running the Application

1. Ensure Go is installed
2. Run the application:
   ```bash
   go run main.go
   ```

## Testing the API

Examples of testing each endpoint:

1. Create a book:
   ```bash
   curl -X POST http://localhost:8080/books \
     -H "Content-Type: application/json" \
     -d '{
       "title": "The Go Programming Language",
       "author": "Alan A. A. Donovan",
       "publishedYear": 2015
     }'
   ```

2. Get all books:
   ```bash
   curl http://localhost:8080/books
   ```

3. Get book by ID:
   ```bash
   curl http://localhost:8080/books/1
   ```

4. Search books by author:
   ```bash
   curl http://localhost:8080/books?author=Alan%20A.%20A.%20Donovan
   ```

5. Delete a book:
   ```bash
   curl -X DELETE http://localhost:8080/books/1
   ```

## Implementation Notes

- Used clean architecture principles with separation of concerns
- Implemented thread-safe operations for concurrent access
- Added comprehensive error handling and input validation
- Included panic recovery for unexpected errors