// okaq gira pasu
// sasu web server
// AQ <aq@okaq.com>
// 2015-10-18
package main

import (
    "fmt"
    "net/http"
)

func RasuHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    http.ServeFile(w,r,"jasu.html")
}

func main() {
    fmt.Println("web server on okaq gira pasu rasu jasu")
    http.HandleFunc("/", RasuHandler)
    http.ListenAndServe(":8008", nil)
}
