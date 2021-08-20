package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	fmt.Println("Starting server...")

	http.HandleFunc("/", hello)

	log.Fatal(http.ListenAndServe(":8000", nil))
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Welcome!")
}
