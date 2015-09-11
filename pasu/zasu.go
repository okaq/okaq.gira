// okaq gira pasu
// zasu web server
// btc scalp trainer
// AQ <aq@okaq.com>
// 2015-09-12
package main

import (
    "fmt"
    "net/http"
)

var (
    B int
    D []string
    T map[string]string
)

func Data() {
    B = 4
    D = make([]string, B)
    D[0] = "wild invetion"
    D[1] = "party out of mind"
    D[2] = "insane root"
    D[3] = "bloody man"
}

func Table() {
    T = make(map[string]string)
    for i := 0; i < B; i++ {
        T[D[i]] = // hash string
    }
}

func main() {
    fmt.Println("starting btc scalp web app!")
    Data()
    Table()
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "basu.html")
    })
    http.ListenAndServe(":8008", nil)
}
