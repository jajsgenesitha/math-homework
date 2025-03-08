package main

import (
    "fmt"
    "math/rand"
)

func main() {
    rand.Seed(42) // initialize seed for random number generator
    x := rand.Intn(10) // generate a random integer between 0 and 10, inclusive
    fmt.Println("Random number:", x)
}
