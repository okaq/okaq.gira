// okaq naka web app
// AQ <aq@okaq.com>
// 2015-08-20
package main

import (
    "net/http"
)

func init() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "dami.html")
    })
}
