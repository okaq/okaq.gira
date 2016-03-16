package main

import (
    "fmt"
    "net/http"
)

const (
    IIKO = "iiko.html"
)

func RikoHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    http.ServeFile(w,r,IIKO)
}

func main() {
    fmt.Println("okaq gira qiko riko web running on localhost:8008")
    http.HandleFunc("/", RikoHandler)
    http.ListenAndServe(":8008", nil)
}
