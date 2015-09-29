// okaq gira pasu
// uasu web server
// AQ <aq@okaq.com>
// 2015-09-30
package main

import (
    "fmt"
    "net/http"
)

func UasuHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    http.ServeFile(w,r,"gasu.html")
}

func main() {
    fmt.Println("web server on okaq gira pasu uasu")
    http.HandleFunc("/", UasuHandler)
    http.ListenAndServe(":8008", nil)
}
