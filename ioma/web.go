// golang web server
// AQ <aq@okaq.com>
// 2015-01-06
package main

import (
    "net/http"
)

const (
    // HTML = "index.html"
    SHIN = "shin.html" // new
)

func SelectorHandler(w http.ResponseWriter, req *http.Request) {
    http.ServeFile(w, req, SHIN)
}

func init() {
    http.HandleFunc("/", SelectorHandler)
}
