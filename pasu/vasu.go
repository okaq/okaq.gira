// okaq gira pasu
// vasu web server
// AQ <aq@okaq.com>
// 2015-09-24
package main

import (
    "fmt"
    "net/http"
)

func VasuHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    http.ServeFile(w,r,"fasu.html")
}

func main() {
    fmt.Println("start web server okaq gira pasu vasu...")
    http.HandleFunc("/", VasuHandler)
    http.ListenAndServe(":8008", nil)
}


