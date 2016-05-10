// golang web server
// okaq partizan nano game
// AQ <aq@okaq.com>
// 2016-05-12
package main

import (
    "net/http"
)

const (
    INDEX = "boki.html"
)

func YokiHandler(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, INDEX)
}

func main() {
    http.HandleFunc("/", YokiHandler)
    http.ListenAndServe(":8008", nil)
}
