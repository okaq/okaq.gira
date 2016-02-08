package main

import (
    "fmt"
    "net/http"
)

const (
    FIKO = "fiko.html"
)

func UikoHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    http.ServeFile(w, r, FIKO)
}

func main() {
    fmt.Println("okaq gir qikko uiko web server started on localhost:8008")
    http.HandleFunc("/", UikoHandler)
    http.ListenAndServe(":8008", nil)
}
