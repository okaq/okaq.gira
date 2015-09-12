// okaq gira pasu
// yasu web server
// AQ <aq@okaq.com>
// 2015-09-14
package main

import (
    "fmt"
    "net/http"
)

func main() {
    fmt.Println("yasu web server live...")
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w,r,"casu.html")
    })
    http.ListenAndServe(":8008", nil)
}
