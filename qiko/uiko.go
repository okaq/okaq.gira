package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
)

const (
    FIKO = "fiko.html"
)

var (
    // Cache []string
    // Rec chan string
    Cache []Qid
    Rec chan Qid
    Sen chan []byte
    Dat chan []byte
)

type Qid []byte

func UikoHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    http.ServeFile(w, r, FIKO)
}

func QidHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    // w.Write([]byte("ok qid"))
    // decode and display json body
    var b0 map[string]byte
    dec := json.NewDecoder(r.Body)
    err := dec.Decode(&b0)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(b0)
    s0 := fmt.Sprintf("qid: rec %d bytes", len(b0))
    w.Write([]byte(s0))
}

func BidHandler(w http.ResponseWriter, r *http.Request) {
    // binary post data handler
    fmt.Println(r)
    // fmt.Println(r.Body)
    // w.Write([]byte("bid ok"))
    // readall for small data streams
    // use bufio for larger unknown streams
    b0, err := ioutil.ReadAll(r.Body)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(b0)
    // pitch bytes on chan
    Sen <-b0
    // each write cascades a sequential read
    // which is sent to the client
    // wait here for the cache to finish writing
    j0 := <-Dat
    fmt.Println(string(j0))
    // s0 := fmt.Sprintf("bid: rec %d bytes", len(b0))
    // w.Write([]byte(s0))
    w.Write(j0)
}

func Catch() {
    // set up global chan and cache
    Cache = []Qid{}
    Rec = make(chan Qid)
    Dat = make(chan []byte)
    go func() {
        // fmt.Println(Cache)
        // fmt.Println(Rec)
        for {
            q0 := <-Rec
            Cache = append(Cache, q0)
            fmt.Println(Cache)
            // send json string
            j0, err := json.Marshal(Cache)
            if err != nil {
                fmt.Println(err)
            }
            fmt.Println(j0)
            Dat <-j0
        }
    }()
}

func Pitch() {
    // recieve bytes and convert to Qid
    Sen = make(chan []byte)
    go func() {
        for {
            b0 := <-Sen
            q0 := Qid(b0)
            Rec <-q0
        }
    }()
}

func main() {
    fmt.Println("okaq gir qikko uiko web server started on localhost:8008")
    // launch goroutine to handle cache
    // receive on requests to append to cache
    // request handler sends on global chan
    Catch()
    // dedicated type for Cache and Receiver
    // json send and request via fetch
    Pitch()
    // each request handled adds
    // type binary id to list
    // list index order determines first in
    http.HandleFunc("/", UikoHandler)
    http.HandleFunc("/qid", QidHandler)
    http.HandleFunc("/bid", BidHandler)
    http.ListenAndServe(":8008", nil)
}
