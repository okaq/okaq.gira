// web server for google code jam 2016 finals
// problem e. radioactive islands
package main

import (
    "bufio"
    "fmt"
    "os"
    "net/http"
    "strconv"
)

const (
    INDEX = "radio.html"
    LARGE = "E-large-practice.in"
    SMALL = "E-small-practice.in"
    // output files for solution
)

type Island struct {
    Y float32 // y-coord
    X float32 // x-coord
}

func NewIsland() *Island {
    I := Island{}
    I.X = 0.0
    return &I
}

func Read() {
    // open files
    // parse data into Island[]
    fi, err := os.Open(SMALL)
    if err != nil {
        fmt.Println(err)
    }
    s := bufio.NewScanner(fi)
    s.Scan()
    t, err := strconv.Atoi(s.Text())
    if err != nil {
        fmt.Println(err)
    }
    fmt.Printf("number of test cases: %d\n", t)
    // stdin, fmt scan
    for i := 0; i < t; i++ {
        s.Scan()
        s0 := s.Text()
        fmt.Println(s0)
        // int float float
        s.Scan()
        s1 := s.Text()
        fmt.Println(s1)
        // float
    }
}

func RadioHandle(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
}

func TestHandle(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    // json test data
}

func Integrate() {
    // pure definite integral numerical simulation
    // are under the function, accurate to 0.001
}

func Linear(x0 float32) float32 {
    // linear slope equals one
    return x0
}

func main() {
    fmt.Println("reading input data")
    Read()
    fmt.Println("starting web server on localhost:8080")
    http.HandleFunc("/", RadioHandle)
    http.HandleFunc("/t", TestHandle)
    http.ListenAndServe(":8080", nil)
}
