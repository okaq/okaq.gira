package main

import (
    "crypto/sha1"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "strconv"
    "time"
)

const (
    // total players
    LIMIT = 128
    // random player sample
    SAMP = 8
)

var (
    // player cache
    Q [LIMIT]*Qid
    // chans to send new player to cache
    // read a random sample
    W chan *Qid
    R chan [SAMP]*Qid
    Index int
)

type Qid struct {
    Time string
    Perf string
    Id string
    Hash [sha1.Size]byte
}

func YikoHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    http.ServeFile(w,r,"biko.html")
}

func PingHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    // w.Write([]byte("ok dues"))
    // qid instance 
    // create new, store in cache
    t0 := time.Now().UnixNano()
    t1 := strconv.FormatInt(t0, 10)
    q0 := Qid{Time:t1}
    b0, err := json.Marshal(q0)
    if err != nil {
        fmt.Println(err)
    }
    w.Write(b0)
}

func PongHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    // w.Write([]byte("ok ello"))
    // set rest of values in cache
    // hash all three markers to qid
    b0, err := ioutil.ReadAll(r.Body)
    if err != nil {
        fmt.Println(err)
    }
    var q0 Qid
    err = json.Unmarshal(b0, &q0)
    if err != nil {
        fmt.Println(err)
    }
    // fmt.Println(q0)
    // h0 := sha1.New()
    h1 := sha1.Sum(b0)
    // fmt.Printf("%x", h1)
    // fmt.Println(string(h1))
    // q0.Hash = string(h1)
    q0.Hash = h1
    fmt.Printf("Qid: Time: %s. Id: %s. Perf: %s. Hash: %x\n", q0.Time, q0.Id, q0.Perf, q0.Hash)
    w.Write(q0.Hash[:])
}

func QingHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
}

func State() {
    W = make(chan *Qid)
    R = make(chan [SAMP]*Qid)
    Index = 0
    go Cache()
    q0 := Qid{Time:"0",Id:"0",Perf:"0"}
    W <- &q0
}

func Cache() {
    for {
        select {
            case w := <-W:
            fmt.Println(w)
            case r := <-R:
            fmt.Println(r)
            // reader contains a chan 
            // which blocks while awaiting sample
            // fill sampling slice with [1..LIMT]
            // remove indicies with
            // a = append(a[:i],a[i+1:])
        }
    }
}

func main() {
    fmt.Println("creating state")
    State()
    fmt.Println("starting web server")
    http.HandleFunc("/", YikoHandler)
    http.HandleFunc("/a", PingHandler)
    http.HandleFunc("/b", PongHandler)
    http.HandleFunc("/c", QingHandler)
    http.ListenAndServe(":8008", nil)
}
