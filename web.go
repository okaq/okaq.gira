// golang web server 
// deployed to google cloud
// AQ <aq@okaq.com>
// 2014-06-21
package main

import (
    "net/http"
)

func init() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "index.html")
    })
}
