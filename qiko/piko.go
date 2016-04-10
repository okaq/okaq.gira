// golang web server
// okaq tsunkeri greet page
// AQ <aq@okaq.com>
// 2016-04-04
package main

import (
    "net/http"
)

const (
    INDEX = "kiko.html"
)

func PikoHandler(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, INDEX)
}

func main() {
    http.HandleFunc("/", PikoHandler)
    http.ListenAndServe(":8008", nil)
}
