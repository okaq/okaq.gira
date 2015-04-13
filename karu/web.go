// golang web server 
// dev env
// AQ <aq@okaq.com>
// 2015-04-04
package main

import (
    "net/http"
    "os"
)

const (
    BADE = "bade.html"
    CAPI = "capi.html"
    ENKO = "enko.html"
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
    File *os.File
    // sync.RWLock
}

func NewCapis() *Capis {
    c0 := new(Capis)
    c0.Name = "dapi.json"
    var err error
    if (c0.Exists()) {
        c0.File, err = os.Open(c0.Name)
        if err != nil {
            panic("cannot open file")
        }
    } else {
        c0.File, err = os.Create(c0.Name)
        if err != nil {
            panic("cannot create file")
            // always panics in app engine
            // distributed filesystem
            // does not allow file creation
        }
    }
    return c0
}

func (c *Capis) Exists() bool {
    if _, err := os.Stat(c.Name); err != nil {
        if os.IsNotExist(err) {
            return false
        }
    }
    return true
}
// use blobstore api for large file upload
// datastore for persistent sql
// memcache for short lived globals memory

// go-fsnotify for file notifications

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
    http.HandleFunc("/enko", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, ENKO)
    })
    http.HandleFunc("/capi/init", CapiHandler)
}

// may initially seed root request handler
// with template that emits server info, token

