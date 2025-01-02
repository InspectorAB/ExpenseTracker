package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Welcome to the Expense Tracker API!")
    })

    log.Println("Server is running on port 8080...")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
