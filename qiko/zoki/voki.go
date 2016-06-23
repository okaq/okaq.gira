// golang web server
// okaq war / partizan nano game
// AQ <aq@okaq.com>
// 2016-06-24
package main

import (
    "fmt"
    "net/http"
)

const (
    INDEX = "eoki.html"
)

func VokiHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    http.ServeFile(w, r, INDEX)
}

func main() {
    http.HandleFunc("/", VokiHandler)
    http.ListenAndServe(":8008", nil)
}
