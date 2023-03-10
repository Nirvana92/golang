package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting the server ....")

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, I love %s!", r.URL.Path[1:])
}
