// golang web server
// gcloud dev env
// AQ <aq@okaq.com>
// 2015-06-22
package main

import (
    "net/http"
)

const (
    TOGU = "togu.html"
)

func init() {
    http.HandleFunc("/togu", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, TOGU)
    })
}
