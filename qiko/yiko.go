package main

import (
    "fmt"
    "net/http"
)

func YikoHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    http.ServeFile(w,r,"biko.html")
}

func main() {
    fmt.Println("starting web server")
    http.HandleFunc("/", YikoHandler)
    http.ListenAndServe(":8008", nil)
}
