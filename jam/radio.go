// web server for google code jam 2016 finals
// problem e. radioactive islands
package main

import (
    "bufio"
    "encoding/json"
    "fmt"
    "os"
    "net/http"
    "strconv"
    "strings"
    "sync"
)

const (
    INDEX = "radio.html"
    LARGE = "E-large-practice.in"
    SMALL = "E-small-practice.in"
    // output files for solution
)

var (
    Islands []*Island
    WG sync.WaitGroup
    // use channels instead
)

type Island struct {
    Y float32 // y-coord
    X float32 // x-coord
    N int // total
    A float32 // start
    B float32 // end
    C []float32 // islands
    // solution
    S []float32
    Id int // case id
}

func NewIsland() *Island {
    I := Island{}
    I.X = 0.0
    I.Y = 0.0
    I.N = 0
    return &I
}

func (I *Island) Solve() {
    I.S = make([]float32, 100)
    WG.Done()
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
    Islands = make([]*Island, t)
    // stdin, fmt scan
    for i := 0; i < t; i++ {
        s.Scan()
        s0 := s.Text()
        // fmt.Println(s0)
        // int float float
        var n0 int
        var x0 float32
        var y0 float32
        fmt.Sscanf(s0, "%d %f %f", &n0, &x0, &y0)
        fmt.Printf("values: n = %d, x = %f, y = %f\n", n0, x0, y0)
        s.Scan()
        s1 := s.Text()
        // fmt.Println(s1)
        // floats
        s2 := strings.Split(s1, " ")
        s3 := make([]float32, n0)
        for j := 0; j < n0; j++ {
            s4, err := strconv.ParseFloat(s2[j], 32)
            if err != nil {
                fmt.Println(err)}
            s3[j] = float32(s4)
        }
        fmt.Println(s2, s3)
        I := NewIsland()
        I.Id = i + 1
        I.N = n0
        I.A = x0
        I.B = y0
        I.C = s3
        fmt.Println(I)
        Islands[i] = I
        WG.Add(1)
        go Islands[i].Solve()
        // prob def - Island
        // index i - test case #
        // boat start coordinates (X,Y)
        // islands (N, []Y)
        // grand game list - []Island
        // test # is index + 1
        // solution is the path function

        // go I.Solve()
        // sync add wait group
    }
}

func RadioHandle(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
}

func TestHandle(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    // json test data
    // send Island[] list
    // client side Island class
    j0, err := json.Marshal(Islands)
    if err != nil {
        fmt.Println(err)
    }
    // fmt.Println(j0)
    w.Write(j0)
}

func Integrate() {
    // pure definite integral numerical simulation
    // are under the function, accurate to 0.001
}

func Linear(x0 float32) float32 {
    // linear slope equals one
    return x0
}

func Square(x0 float32) float32 {
    return x0 * x0
}

func Sum() {
    a0 := float32(0.0)
    for i := 0; i < 1000000; i++ {
        x0 := float32(i) * 0.000001
        y0 := Square(x0) * 0.000001
        // y0 := Linear(x0) * 0.000001
        a0 = a0 + y0
    }
    fmt.Printf("Definite integral = %f\n", a0)
    // 0.333299, accurate to 4 digits
}

func main() {
    fmt.Println("reading input data")
    Read()
    fmt.Println("starting web server on localhost:8080")
    Sum()
    WG.Wait()
    fmt.Println(Islands[0].S)
    http.HandleFunc("/", RadioHandle)
    http.HandleFunc("/t", TestHandle)
    http.ListenAndServe(":8080", nil)
}
