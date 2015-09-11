// okaq gira pasu
// zasu web server
// btc scalp trainer
// AQ <aq@okaq.com>
// 2015-09-12
package main

import (
    "crypto/sha256"
    "encoding/json"
    "fmt"
    "hash"
    "net/http"
)

var (
    B int // size of data
    D []string // initial string data
    T map[string]string // hash table
    U map[string][]byte // byte wise hash table
    H hash.Hash // sha256 hash func
    J []byte // string json encoding
    K []byte // byte json encoding
)

func Data() {
    B = 4
    D = make([]string, B)
    D[0] = "wild invetion"
    D[1] = "party out of mind"
    D[2] = "insane root"
    D[3] = "bloody man"
    H = sha256.New()
}

func Table() {
    T = make(map[string]string)
    U = make(map[string][]byte)
    for i := 0; i < B; i++ {
        H.Reset()
        b0 := []byte(D[i])
        H.Write(b0)
        b1 := H.Sum(nil)
        T[D[i]] = string(b1)
        U[D[i]] = b1
    }
    fmt.Println(T)
    // json strings valid UTF-8
    var err error
    J, err = json.Marshal(T)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(string(J))
    K, err = json.Marshal(U)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(U)
}

func ZasuHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    http.ServeFile(w, r, "basu.html")
}

func AasuHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    w.Write([]byte("ok aasu"))
}

func BasuHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    w.Write([]byte("ok basu"))
}

func main() {
    fmt.Println("starting btc scalp web app!")
    Data()
    Table()
    // http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
       //  http.ServeFile(w, r, "basu.html")
    // })
    http.HandleFunc("/", ZasuHandler)
    http.HandleFunc("/a", AasuHandler)
    http.HandleFunc("/b", BasuHandler)
    http.ListenAndServe(":8008", nil)
}
