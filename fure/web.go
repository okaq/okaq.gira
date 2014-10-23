// golang web server
// deployed to GCloud App Engine (golang)
// AQ <aq@okaq.com>
// 2014-10-16
package main

import (
    "net/http"
)

func init() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { http.ServeFile(w,r,"index.html") })
}
