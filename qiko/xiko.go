package main

import (
    "fmt"
    "net/http"
)

func XikoHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    http.ServeFile(w,r,"ciko.html")
}

func main() {
    fmt.Println("starting okaq gira qiko xiko web server on http://localhost:8008/")
    http.HandleFunc("/", XikoHandler)
    http.ListenAndServe(":8008", nil)
}
