// golang web server 
// dev env
// AQ <aq@okaq.com>
// 2015-04-04
package main

import (
    "net/http"
)

const (
    BADE = "bade.html"
)

func init() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, BADE)
    })
    http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("test ok if you can read this"))
    })
}

// may initially seed root request handler
// with template that emits server info, token

