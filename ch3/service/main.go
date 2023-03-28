package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Server is running on http://localhost:8080")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to my website!")
	})

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "About me!")
	})

	http.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Contact me!")
	})

	http.ListenAndServe(":8080", nil)
}
