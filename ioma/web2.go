// okaq gira ioma nako
// bmp ping pong
// golang web server
// AQ <aq@okaq.com>
// 2015-02-10
package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "net/http"
    "sync"
)

const (
    NAKO = "nako.html" // bmp ping pong
    OTOH = "otoh.html" // 1fps binary bitmap anim
    RUKI = "ruki.html" // bezier curve alpha
)

var (
    P *Peers
)

type Peers struct {
    Ids []string
    Users int
    Index int
    sync.Mutex
}

func NewPeers() *Peers {
    p := new(Peers)
    p.Users = 8
    p.Index = 0
    p.Ids = make([]string, p.Users)
    return p
}

func (p *Peers) Add(s0 string) {
    p.Lock()
    defer p.Unlock()
    if p.Index <= (p.Users - 1) {
        p.Ids[p.Index] = s0
        p.Index = p.Index + 1
    }
}

func (p *Peers) Json() ([]byte, error) {
    p.Lock()
    defer p.Unlock()
    b0, err := json.Marshal(p)
    if err != nil {
        return nil, err
    }
    return b0, nil
}

// root static html handle
// single client html5 web app for xhr, net, conn
func RootHandler(w http.ResponseWriter, req *http.Request) {
    http.ServeFile(w, req, RUKI)
}

// peers cache ajax response
// request POSTs id obtained from peerjs
func PeersHandler(w http.ResponseWriter, req *http.Request) {
    b0 := new(bytes.Buffer)
    b0.ReadFrom(req.Body)
    s0 := string(b0.Bytes())
    // use chan here to communicate data
    P.Add(s0)
    b1, err := P.Json()
    if err != nil {
        w.Header().Set("Content-Type", "text/plain")
        s1 := fmt.Sprintf("error: %s\n", err)
        w.Write([]byte(s1))
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.Write(b1)
    // w.Write([]byte("ok!"))
}

func init() {
    P = NewPeers()
    http.HandleFunc("/", RootHandler)
    http.HandleFunc("/peers", PeersHandler)
}
