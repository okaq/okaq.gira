// golang web server
// okaq war / partizan nano game
// AQ <aq@okaq.com>
// 2016-06-16
package main

import (
    "fmt"
    "net/http"
)

const (
    INDEX = "doki.html"
)

func WokiHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    http.ServeFile(w, r, INDEX)
}

func main() {
    http.HandleFunc("/", WokiHandler)
    http.ListenAndServe(":8008", nil)
}
