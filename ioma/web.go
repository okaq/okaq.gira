// golang web server
// AQ <aq@okaq.com>
// 2015-01-06
package main

import (
    "encoding/json"
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
    APIKEY = "mrldpk6ryr0e8kt9"
)

var (
    C *Counter
    PC *PeerCache
)

type Counter struct {
    Id int
    Conn bool
    sync.Mutex
}

func NewCounter() *Counter {
    c0 := new(Counter)
    c0.Id = 0
    c0.Conn = true
    return c0
}

func (c *Counter) Increment() {
    c.Lock()
    defer c.Unlock()
    c.Id = c.Id + 1
    c.Conn = !c.Conn
}

type Peer struct {
    Id int
    Key string
    Conn bool
}

type PeerCache struct {
    Ids []string
    Index int
    Users int
    sync.Mutex
}

func NewPeerCache() *PeerCache {
    pc0 := new(PeerCache)
    pc0.Users = 10
    pc0.Ids = make([]string, pc0.Users)
    pc0.Index = 0
    return pc0
}

func (pc *PeerCache) Add(s0 string) {
    pc.Lock()
    defer pc.Unlock()
    pc.Ids[pc.Index] = s0
    if (pc.Index == pc.Users - 1) {
        pc.Index = pc.Index + 1
    } else {
        pc.Index = 0
    }
}

// struct RWLock peerid counter, cycles from (0-31)
// client can merely try to connect to [id-1,id+1]
// json struct, {apiKey,peerId}

// cache of peer ids
// represented by a slice and a pointer
// each game request, generate id (increment counter)
// store in cache at pointer location, advance pointer
// return json object with cache
// cycle when pointer reaches end, overwriting

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

func PeerHandler(w http.ResponseWriter, req *http.Request) {
}

func KeyHandler(w http.ResponseWriter, req *http.Request) {
    C.Increment()
    p0 := Peer{
        Id: C.Id,
        Key: APIKEY,
        Conn: C.Conn,
    }
    b0, err := json.Marshal(p0)
    if err != nil {
        s0 := fmt.Sprintf("error: %s\n", err)
        w.Write([]byte(s0))
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.Write(b0)
}

func PeerJSHandler(w http.ResponseWriter, req *http.Request) {
    http.ServeFile(w, req, PEERJS)
}

func init() {
    C = NewCounter()
    PC = NewPeerCache()
    http.HandleFunc("/", SelectorHandler)
    http.HandleFunc("/count", CountHandler)
    http.HandleFunc("/key", KeyHandler)
    http.HandleFunc("/peer", PeerHandler)
    http.HandleFunc("/peer.min.js", PeerJSHandler)
}
