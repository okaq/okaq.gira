// golang web server
// AQ <aq@okaq.com>
// 2014-12-14
package main

import (
    "net/http"
)

func init() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { http.ServeFile(w,r,"okaq.html") })
    http.HandleFunc("/kk", func(w http.ResponseWriter, r *http.Request) { http.ServeFile(w,r,"kk.html") })
    http.HandleFunc("/mj", func(w http.ResponseWriter, r *http.Request) { http.ServeFile(w,r,"mj.html") })
}
