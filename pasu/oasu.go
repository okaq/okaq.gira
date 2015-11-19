// okaq gira oasu
// pasu web server
// AQ <aq@okaq.com>
// 2015-11-12
package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
)

var (
    R chan *Ro
    W chan *Wo
    Q []*Qid
)

type Qid struct {
    Time int
    Id int
}

// read op
type Ro struct {
    // i int
    r chan []*Qid
}

// write op
type Wo struct {
    // i int
    Q *Qid
    r chan bool
}

func OasuHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    http.ServeFile(w,r,"masu.html")
    // stateful goroutine pattern
    // to cache reads and writes of player ids
}

func UsoaHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    fmt.Println("req body")
    fmt.Println(r.Body)
    b0, err := ioutil.ReadAll(r.Body)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println("raw data")
    fmt.Println(b0)
    fmt.Println("string representation")
    fmt.Println(string(b0))
    fmt.Println("json data")
    var q0 Qid
    err = json.Unmarshal(b0, &q0)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(q0)
    w.Write([]byte("ok xhr!"))
}

func SoauHandler(w http.ResponseWriter, r *http.Request) {
    // write qid to cache
    b0, err := ioutil.ReadAll(r.Body)
    if err != nil {
        fmt.Println(err)
    }
    var q0 Qid
    err = json.Unmarshal(b0, &q0)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(q0)
    // send new Writer on global chan
    w0 := &Wo{
        Q: &q0,
        r: make(chan bool)}
    W <- w0
    w.Write(b0)
}

func AusoHandler(w http.ResponseWriter, r *http.Request) {
    // read quid list and write to response
    w.Write([]byte("ok auso!"))
}

func State() {
    // comm channels
    R = make(chan *Ro)
    W = make(chan *Wo)
    // central state
    Q = make([]*Qid, 1)
    // ready player zero
    Q[0] = new(Qid)
    Q[0].Time = 0
    Q[0].Id = 0
    // launch goroutines
    // selector
    go func() {
        for {
            select {
                case r := <-R:
                fmt.Println(r)
                case w := <-W:
                fmt.Println(w)
            }
        }
    }()
    // reader - stats request handler
    // writer - player id post handler
}

func main() {
    fmt.Println("web server on okaq gira pasu oasu masu")
    fmt.Println("starting state")
    State()
    fmt.Println("assigning web handlers")
    http.HandleFunc("/", OasuHandler)
    http.HandleFunc("/a", UsoaHandler)
    http.HandleFunc("/b", SoauHandler)
    http.HandleFunc("/c", AusoHandler)
    fmt.Println("web server running...")
    http.ListenAndServe(":8008", nil)
}


