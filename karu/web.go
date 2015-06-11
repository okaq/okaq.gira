// golang web server
// gcloud dev env
// AQ <aq@okaq.com>
// 2015-06-12
package main

import (
    "net/http"
)

const (
    ONKU = "onku.html"
)

func init() {
    http.HandleFunc("/onku", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, ONKU)
    })
}
