package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    var n = 100
    rand.Seed(time.Now().UnixNano())

    for i := 0; i < n; i++ {
        r := math.Rand()
        if r <= 0.5 {
            fmt.Printf("True\n")
        } else {
            fmt.Printf("False\n")
        }
    }

}
