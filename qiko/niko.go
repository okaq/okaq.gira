// golang web server
// okaq.com home, line transform
// AQ <aq@okaq.com>
// 2016-05-02
package main

import (
    "net/http"
)

const (
    INDEX = "miko.html"
)

func NikoHandler(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, INDEX)
}

func main() {
    http.HandleFunc("/", NikoHandler)
    http.ListenAndServe(":8008", nil)
}
