// okaq gira pasu
// aasu web server
// AQ <aq@okaq.com>
// 2015-10-10
package main

import (
    "fmt"
    "net/http"
)

func SasuHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    http.ServeFile(w,r,"iasu.html")
}

func main() {
    fmt.Println("web server on okaq gira pasu sasu iasu")
    http.HandleFunc("/", SasuHandler)
    http.ListenAndServe(":8008", nil)
}
