package main 

import (
    "net/http"
    "fmt"
)

// Default Request Handler
func defaultHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Meeli Vyas")
}

func main1() {
    http.HandleFunc("/", defaultHandler)
    http.ListenAndServe(":8080", nil)
}

