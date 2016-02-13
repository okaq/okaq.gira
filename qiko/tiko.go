package main

import (
    "fmt"
    "net/http"
)

const (
    GIKO = "giko.html"
)

func TikoHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    http.ServeFile(w, r, GIKO)
}

func main() {
    fmt.Println("okaq gira qiko web server on localhost:8008")
    http.HandleFunc("/", TikoHandler)
    http.ListenAndServe(":8008", nil)
}
