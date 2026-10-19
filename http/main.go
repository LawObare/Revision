package main

import (
	"fmt"
	"net/http"
)

func main() {

	


	http.HandleFunc("/about", about)
	http.HandleFunc("/", home)


	fmt.Println("Listening on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world"))
}

func about(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is my first real server built by hand, just as practice!!!"))
}