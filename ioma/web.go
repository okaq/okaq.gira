// golang web server
// AQ <aq@okaq.com>
// 2015-01-06
package main

import (
    "net/http"
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
)

func SelectorHandler(w http.ResponseWriter, req *http.Request) {
    http.ServeFile(w, req, YOMI)
}

func PeerJSHandler(w http.ResponseWriter, req *http.Request) {
    http.ServeFile(w, req, PEERJS)
}

func init() {
    http.HandleFunc("/", SelectorHandler)
    http.HandleFunc("/peer.min.js", PeerJSHandler)
}
