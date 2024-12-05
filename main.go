package main

import (
    "net/http"
    "github.com/go-chi/chi/v5"
)

func main() {
    // Initialise un nouveau routeur Chi
    r := chi.NewRouter()

	http.ListenAndServe(":8080", r)
}