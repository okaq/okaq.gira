// golang web server for Google Cloud (GAE)
// okaq gira aka nano game
// AQ <aq@okaq.com>
// 2014-05-30
package main

import (
    "net/http"
)

const (
    PATH = "index.html"
)

func init() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, PATH)
    })
}
