package main

import (
	"fmt"
	"net/http"
)

func main() {
	// For local testing, we use a simple server
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Running locally")
	})
	
	fmt.Println("Local server running on :8080")
	http.ListenAndServe(":8080", nil)
}