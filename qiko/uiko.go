package main

import (
    "fmt"
    "net/http"
)

const (
    FIKO = "fiko.html"
)

var (
    Cache []string
    Rec chan string
)

func UikoHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    http.ServeFile(w, r, FIKO)
}

func main() {
    fmt.Println("okaq gir qikko uiko web server started on localhost:8008")
    // launch goroutine to handle cache
    // receive on requests to append to cache
    // request handler sends on global chan
    
    // dedicated type for Cache and Receiver
    // json send and request via fetch
    http.HandleFunc("/", UikoHandler)
    http.ListenAndServe(":8008", nil)
}
