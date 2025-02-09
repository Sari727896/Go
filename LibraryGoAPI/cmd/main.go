package main

import (
	"library-api/internal/routes"
	"log"
	"net/http"
)

func main() {

    // Initialize router
    r := router.SetupRouter()

    // Start server
    log.Println("Server starting on port 8080...")
    if err := http.ListenAndServe(":8080", r); err != nil {
        log.Fatal(err)
    }
}