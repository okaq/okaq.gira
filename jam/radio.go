// web server for google code jam 2016 finals
// problem e. radioactive islands
package main

import (
    "fmt"
    "net/http"
)

const (
    INDEX = "radio.html"
)

func HandleFunc(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
}

func main() {
    fmt.Println("starting web server on localhost:8080")
    http.HandleFunc("/", RadioHandle)
    http.ListenAndServe(":8080", nil)
}
