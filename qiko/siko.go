package main

import (
    "fmt"
    "net/http"
)

const (
    HIKO = "hiko.html"
)

func SikoHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    http.ServeFile(w,r,HIKO)
}

func main() {
    fmt.Println("okaq gira qiko siko web running on localhost:8008")
    http.HandleFunc("/", SikoHandler)
    http.ListenAndServe(":8008", nil)
}
