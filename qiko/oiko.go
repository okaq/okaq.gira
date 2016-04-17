// golang web server
// one2.one web dictaphone
// AQ <aq@okaq.com>
// 2016-04-18
package main

import (
    "net/http"
)

const (
    INDEX = "liko.html"
)

func OikoHandler(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, INDEX)
}

func main() {
    http.HandleFunc("/", OikoHandler)
    http.ListenAndServe(":8008", nil)
}
