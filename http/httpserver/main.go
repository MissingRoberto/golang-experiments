package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!")
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal("ListenAndServe:", http.ListenAndServe("localhost:8080", nil))
}
