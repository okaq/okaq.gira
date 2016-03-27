package main

import (
    "fmt"
    "net/http"
)

const (
    JIKO = "jiko.html"
)

func QikoHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    http.ServeFile(w,r,JIKO)
}

func main() {
    fmt.Println("okaq gira qiko qiko web running on localhost:8008")
    http.HandleFunc("/", QikoHandler)
    http.ListenAndServe(":8008", nil)
}
