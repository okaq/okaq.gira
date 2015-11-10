// okaq gira oasu
// pasu web server
// AQ <aq@okaq.com>
// 2015-11-12
package main

import (
    "fmt"
    "net/http"
)

func OasuHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    http.ServeFile(w,r,"masu.html")
}

func main() {
    fmt.Println("web server on okaq gira pasu oasu masu")
    http.HandleFunc("/", OasuHandler)
    http.ListenAndServe(":8008", nil)
}


