package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "strconv"
    "time"
)

type Qid struct {
    Time string
    Perf string
    Id string
    Hash string
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
    w.Write([]byte("ok ello"))
    // set rest of values in cache
    // hash all three markers to qid
}

func main() {
    fmt.Println("starting web server")
    http.HandleFunc("/", YikoHandler)
    http.HandleFunc("/a", PingHandler)
    http.HandleFunc("/b", PongHandler)
    http.ListenAndServe(":8008", nil)
}
