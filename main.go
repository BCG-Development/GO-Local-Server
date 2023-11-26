package main

import (
    "fmt"
    "net/http"
)

func main() {
    // Define a handler function
    helloHandler := func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, this is your local server!")
    }

    // Register the handler for a specific route
    http.HandleFunc("/", helloHandler)

    // Start the server on port 8080
    fmt.Println("Server running on http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}