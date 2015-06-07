// golang web server
// gcloud dev env
// AQ <aq@okaq.com>
// 2015-06-08
package main

import (
    "net/http"
)

const (
    NONU = "nonu.html"
)

func DapiHandler(w http.ResponseWriter, req *http.Request) {
    w.Write([]byte("greetings xhr"))
}

func init() {
    http.HandleFunc("/nonu", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, NONU)
    })
    http.HandleFunc("/dapi", DapiHandler)
}
// connection / hub message pattern
// channels to broadcast and receive
// xhr connection to post peer id
// server side event push
// once 4 peers have joined
// then peers negotiate random connections
// with one local hub
