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
    CAPI = "capi.html"
)

// capi store
type Capis struct {
    // filename string
    // *os.File
    // sync.RWLock
}

func CapiHandler(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, CAPI)
    // open file (rwlock'd)
    // save init state
}

func init() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, BADE)
    })
    http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("okaq gira karu bade ok!"))
    })
    http.HandleFunc("/logs", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("req logged"))
    })
    http.HandleFunc("/capi", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, CAPI)
    })
    http.HandleFunc("/capi/init", CapiHandler)
}

// may initially seed root request handler
// with template that emits server info, token

