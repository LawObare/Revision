package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/about", about)
	mux.HandleFunc("/", home)

	fmt.Println("Listening on http://localhost:8080")
	http.ListenAndServe(":8080", mux)
}

func home(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	_, err := w.Write([]byte("Hello, world"))
	if err != nil {
		fmt.Errorf("There is an error: %w", err)
	}
}

func about(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is my first real server built by hand, just as practice!!!"))
}
