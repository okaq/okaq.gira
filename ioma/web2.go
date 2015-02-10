// okaq gira ioma nako
// bmp ping pong
// golang web server
// AQ <aq@okaq.com>
// 2015-02-10
package main

import (
    "net/http"
)

const (
    NAKO = "nako.html" // bmp ping pong
)

// root static html handle
// single client html5 web app for xhr, net, conn
func RootHandler(w http.ResponseWriter, req *http.Request) {
    http.ServeFile(w, req, NAKO)
}

// peers cache ajax response
// request POSTs id obtained from peerjs
func PeersHandler(w http.ResponseWriter, req *http.Request) {
    w.Write([]byte("ok!"))
}

func init() {
    http.HandleFunc("/", RootHandler)
    http.HandleFunc("/peers", PeersHandler)
}
