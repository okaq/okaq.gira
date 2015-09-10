// okaq gira pasu
// zasu web server
// btc scalp trainer
// AQ <aq@okaq.com>
// 2015-09-12
package main

import (
    "fmt"
    "net/http"
)

func main() {
    fmt.Println("starting btc scalp web app!")
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "basu.html")
    })
    http.ListenAndServe(":8008", nil)
}
