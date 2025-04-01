package main

import (
    "fmt"
)

func main() {
    // Example of using Go to calculate factorial of numbers from 1 to 5
    fmt.Println("Factorial of numbers from 1 to 5:")
    for i := 1; i <= 5; i++ {
        result := 1
        for j := 2; j <= i; j++ {
            result *= j
        }
        fmt.Printf("%d! = %d\n", i, result)
    }
}
