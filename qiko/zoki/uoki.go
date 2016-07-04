// golang web server
// okaq war / partizan nano game
// AQ <aq@okaq.com>
// 2016-07-06
package main

import (
    "fmt"
    "net/http"
)

const (
    INDEX = "foki.html"
)

func UokiHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    http.ServeFile(w, r, INDEX)
}

func main() {
    http.HandleFunc("/", UokiHandler)
    http.ListenAndServe(":8008", nil)
}
