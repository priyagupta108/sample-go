package main

import (
    "fmt"
    "net/http"

    "github.com/gorilla/mux"
    "github.com/google/uuid"
)

func main() {
    r := mux.NewRouter()

    r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        id := uuid.NewString()
        fmt.Fprintf(w, "Hello, Web! UUID: %s", id)
    })
    http.Handle("/", r)
    http.ListenAndServe(":8080", nil)
}