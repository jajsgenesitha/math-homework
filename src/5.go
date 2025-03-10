package main

import "math/rand"

func main() {
	// Generate a random number between 1 and 100
	randomNumber := rand.Intn(100) + 1

	// Print the random number to the console
	fmt.Println("The random number is", randomNumber)
}
