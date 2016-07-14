// golang web server
// okaq war / partizan nano game
// AQ <aq@okaq.com>
// 2016-07-16
package main

import (
    "fmt"
    "net/http"
)

const (
    INDEX = "goki.html"
)

func TokiHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    http.ServeFile(w, r, INDEX)
}

func main() {
    http.HandleFunc("/", TokiHandler)
    http.ListenAndServe(":8008", nil)
}
