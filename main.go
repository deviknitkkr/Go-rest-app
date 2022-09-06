package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Pong")
	})
	fmt.Println("Starting server...")
	http.ListenAndServe(":8080", nil)
}
