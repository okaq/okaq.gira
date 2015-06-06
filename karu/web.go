// golang web server
// gcloud dev env
// AQ <aq@okaq.com>
// 2015-06-06
package main

import (
    "net/http"
)

const (
    MONU = "monu.html"
)

func DapiHandler(w http.ResponseWriter, req *http.Request) {
    w.Write([]byte("greetings xhr"))
}

func init() {
    http.HandleFunc("/monu", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, MONU)
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
