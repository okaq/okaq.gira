// golang web server
// AQ <aq@okaq.com>
// 2015-01-06
package main

import (
    "net/http"
)

const (
    HTML = "index.html"
)

func SelectorHandler(w http.ResponseWriter, req *http.Request) {
    http.ServeFile(w, req, HTML)
}

func init() {
    http.HandleFunc("/", SelectorHandler)
}
