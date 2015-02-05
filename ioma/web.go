// golang web server
// AQ <aq@okaq.com>
// 2015-01-06
package main

import (
    "fmt"
    "net/http"
    "sync"
)

const (
    // HTML = "index.html"
    // SHIN = "shin.html" // new
    // TETO = "teto.html" // 1fps diffusion
    // UMAO = "umao.html" // voronoise
    // VANI = "vani.html" // peerjs test
    // WOKU = "woku.html" // bitmap oxy font layout and render
    YOMI = "yomi.html" // server side peer id distribution
    PEERJS = "peer.min.js" // PeerJS
    // apiKey
)

var (
    C *Counter
)

type Counter struct {
    Id int
    sync.Mutex
}

func NewCounter() *Counter {
    C := new(Counter)
    C.Id = 0
    return C
}

func (c *Counter) Increment() {
    c.Lock()
    defer c.Unlock()
    c.Id = c.Id + 1
}

// struct RWLock peerid counter, cycles from (0-31)
// client can merely try to connect to [id-1,id+1]
// json struct, {apiKey,peerId}

func SelectorHandler(w http.ResponseWriter, req *http.Request) {
    // C.Increment()
    // fmt.Printf("Web Count: %d\n", C.Id)
    http.ServeFile(w, req, YOMI)
}

func CountHandler(w http.ResponseWriter, req *http.Request) {
    C.Increment()
    s0 := fmt.Sprintf("Web Count: %d\n", C.Id)
    w.Write([]byte(s0))
}

func PeerJSHandler(w http.ResponseWriter, req *http.Request) {
    http.ServeFile(w, req, PEERJS)
}

func init() {
    C = NewCounter()
    http.HandleFunc("/", SelectorHandler)
    http.HandleFunc("/count", CountHandler);
    http.HandleFunc("/peer.min.js", PeerJSHandler)
}
