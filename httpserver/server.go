package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)

	fmt.Println("Starting server on :8080...")
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, Go HTTP server!")
}
