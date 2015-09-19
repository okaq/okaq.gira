// okaq gira pasu
// wasu web server
// AQ <aq@okaq.com>
// 2015-09-20
package main

import (
    "fmt"
    "net/http"
)

func WasuHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    http.ServeFile(w,r,"easu.html")
}

func main() {
    fmt.Println("start web server okaq gira pasu wasu...")
    http.HandleFunc("/", WasuHandler)
    http.ListenAndServe(":8008", nil)
}
