package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("hello world")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "this is the backend")
	})

	http.HandleFunc("/api/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello from go backend")
	})

	fmt.Println("go backend running at port 8000")
	http.ListenAndServe(":8000", nil)
}
