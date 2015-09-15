// okaq gira pasu
// yasu web server
// AQ <aq@okaq.com>
// 2015-09-14
package main

import (
    "fmt"
    "net/http"
)

func XasuHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    http.ServeFile(w,r,"dasu.html")
}

func main() {
    fmt.Println("start web server okaq gira pasu xasu...")
    http.HandleFunc("/", XasuHandler)
    http.ListenAndServe(":8008", nil)
}
