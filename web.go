// golang web server for google cloud
// okaq gira bibi plasma
// AQ <aq@okaq.com>
// 2014-06-03
package main

import (
    "net/http"
)

func init() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "index.html")
    })
}
