// okaq gira pasu
// qasu web server
// AQ <aq@okaq.com>
// 2015-10-20
package main

import (
    "fmt"
    "net/http"
)

func QasuHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    http.ServeFile(w,r,"kasu.html")
}

func main() {
    fmt.Println("web server on okaq gira pasu qasu kasu")
    http.HandleFunc("/", QasuHandler)
    http.ListenAndServe(":8008", nil)
}
