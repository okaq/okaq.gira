// okaq gira oasu
// pasu web server
// AQ <aq@okaq.com>
// 2015-11-12
package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
)

type Qid struct {
    Time int
    Id int
}

func OasuHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    http.ServeFile(w,r,"masu.html")
}

func UsoaHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    fmt.Println("req body")
    fmt.Println(r.Body)
    b0, err := ioutil.ReadAll(r.Body)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println("raw data")
    fmt.Println(b0)
    fmt.Println("string representation")
    fmt.Println(string(b0))
    fmt.Println("json data")
    var q0 Qid
    err = json.Unmarshal(b0, &q0)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(q0)
    w.Write([]byte("ok xhr!"))
}

func main() {
    fmt.Println("web server on okaq gira pasu oasu masu")
    http.HandleFunc("/", OasuHandler)
    http.HandleFunc("/a", UsoaHandler)
    http.ListenAndServe(":8008", nil)
}


