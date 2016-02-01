// golang web server
// okaq arosa greet page
// AQ <aq@okaq.com>
// 2016-02-02
package main

import (
    "net/http"
)

const (
    INDEX = "eiko.html"
)

func VikoHandler(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, INDEX)
}

func main() {
    http.HandleFunc("/", VikoHandler)
    http.ListenAndServe(":8008", nil)
}
