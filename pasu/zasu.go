// okaq gira pasu
// zasu web server
// btc scalp trainer
// AQ <aq@okaq.com>
// 2015-09-12
package main

import (
    "crypto/sha256"
    "fmt"
    "hash"
    "net/http"
)

var (
    B int
    D []string
    T map[string]string
    H hash.Hash
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
    for i := 0; i < B; i++ {
        H.Reset()
        b0 := []byte(D[i])
        H.Write(b0)
        b1 := H.Sum(nil)
        T[D[i]] = string(b1)
    }
    fmt.Println(T)
}

func ZasuHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    http.ServeFile(w, r, "basu.html")
}

func main() {
    fmt.Println("starting btc scalp web app!")
    Data()
    Table()
    // http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
       //  http.ServeFile(w, r, "basu.html")
    // })
    http.HandleFunc("/", ZasuHandler)
    http.ListenAndServe(":8008", nil)
}
