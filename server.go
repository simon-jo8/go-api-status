package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/status" {
        http.Error(w, "404 not found.", http.StatusNotFound)
        return
    }

	fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
}

func main() {
	http.HandleFunc("/", handler)

	fmt.Println("Server is running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}