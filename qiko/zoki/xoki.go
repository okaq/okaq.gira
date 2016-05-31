// golang web server
// okaq war / partizan nano game
// AQ <aq@okaq.com>
// 2016-06-02
package main

import (
    "net/http"
)

const (
    INDEX = "coki.html"
)

func XokiHandler(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, INDEX)
}

func main() {
    http.HandleFunc("/", XokiHandler)
    http.ListenAndServe(":8008", nil)
}
