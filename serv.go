package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Define a handler function for your server
	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, this is a simple web server!")
	}

	// Register the handler function to a route
	http.HandleFunc("/", handler)

	// Start the server on port 8080
	fmt.Println("Server is listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}
