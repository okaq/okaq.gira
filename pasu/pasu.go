// okaq gira pasu
// pasu web server
// AQ <aq@okaq.com>
// 2015-10-30
package main

import (
    "fmt"
    "net/http"
)

func PasuHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    http.ServeFile(w,r,"lasu.html")
}

func main() {
    fmt.Println("web server on okaq gira pasu pasu lasu")
    http.HandleFunc("/", PasuHandler)
    http.ListenAndServe(":8008", nil)
}


