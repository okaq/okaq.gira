// okaq gira pasu
// tasu web server
// AQ <aq@okaq.com>
// 2015-10-06
package main

import (
    "fmt"
    "net/http"
)

func TasuHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    http.ServeFile(w,r,"hasu.html")
}

func main() {
    fmt.Println("web server on okaq gira pasu tasu hasu")
    http.HandleFunc("/", TasuHandler)
    http.ListenAndServe(":8008", nil)
}
