package main

import (
    "fmt"
    "net/http"
)

const (
    INDEX = "diko.html"
)

func WikoHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    http.ServeFile(w,r,INDEX)
}

func main() {
    fmt.Println("starting okaq gira qiko wiko web server on http://localhost:8008/")
    http.HandleFunc("/", WikoHandler)
    http.ListenAndServe(":8008", nil)
}
