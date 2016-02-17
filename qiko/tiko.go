package main

import (
    "fmt"
    "net/http"
)

const (
    GIKO = "giko.html"
)

var (
    // pid cache, first in
    // data chan to send json response list
    // peer id list is json encoded
    // slice of last X pids
    Pids []string
    Pec chan string
)

func TikoHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    http.ServeFile(w, r, GIKO)
}

func PidsCache() {
    Pids = []string{}
    Pec = make(chan string)
    go func() {
        for {
            s0 := <-Pec
            Pids = append(Pids, s0)
            fmt.Println(Pids)
        }
    }()
}

func main() {
    fmt.Println("okaq gira qiko web server on localhost:8008")
    PidsCache()
    http.HandleFunc("/", TikoHandler)
    http.ListenAndServe(":8008", nil)
}
// peerid cache handler
// atomic writes to linked list
// that represents chrono order 
// addition of peers to network
