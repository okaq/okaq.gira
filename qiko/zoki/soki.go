// golang web server
// okaq war / partizan nano game
// AQ <aq@okaq.com>
// 2016-07-28
package main

import (
    "fmt"
    "net/http"
)

const (
    INDEX = "hoki.html"
)

func SokiHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    http.ServeFile(w, r, INDEX)
}

func CountHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
}

func main() {
    http.HandleFunc("/", SokiHandler)
    http.HandleFunc("/c", CountHandler)
    http.ListenAndServe(":8008", nil)
}
