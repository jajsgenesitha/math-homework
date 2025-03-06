package main

import "math/rand"

func main() {
	n := 10
	for i := 0; i < n; i++ {
		fmt.Println(rand.Intn(10))
	}
}
