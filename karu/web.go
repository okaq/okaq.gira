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

var (
    CS *Capis
)

// capi store
// if file not found create it
// load it on first handle
// write on change of state
// format : streaming json encoding
type Capis struct {
    // filename string
    Name string
    // *os.File
    // sync.RWLock
}

func NewCapis() *Capis {
    c0 := new(Capis)
    c0.Name = "capic.json"
    return c0
}

func CapiHandler(w http.ResponseWriter, r *http.Request) {
    // http.ServeFile(w, r, CAPI)
    // open file (rwlock'd)
    // save init state
    w.Write([]byte(CS.Name))
}

func init() {
    CS = NewCapis()
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

