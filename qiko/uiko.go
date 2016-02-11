package main

import (
    "encoding/json"
    "fmt"
    "net/http"
)

const (
    FIKO = "fiko.html"
)

var (
    Cache []string
    Rec chan string
)

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


func main() {
    fmt.Println("okaq gir qikko uiko web server started on localhost:8008")
    // launch goroutine to handle cache
    // receive on requests to append to cache
    // request handler sends on global chan
    
    // dedicated type for Cache and Receiver
    // json send and request via fetch
    http.HandleFunc("/", UikoHandler)
    http.HandleFunc("/qid", QidHandler)
    http.ListenAndServe(":8008", nil)
}
