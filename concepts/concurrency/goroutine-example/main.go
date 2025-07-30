package main

import (
	"fmt"
	"sync"
)

func printSomething(s string, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement the counter when the goroutine completes
	fmt.Println(s)
}

func main() {
	var wg sync.WaitGroup

	words := []string{"alpha", "beta", "gamma", "delta", "phi", "zeta", "eta", "theta", "epsilon"}
	wg.Add(len(words)) // Set the number of goroutines to wait for

	for i, x := range words {
		s := fmt.Sprintf("%d: %s", i+1, x)
		go printSomething(s, &wg) // Start a goroutine for each word
	}

	wg.Wait() // wait for all goroutines to finish
	wg.Add(1) // Add one more for the second print

	printSomething("This is the second thing to be printed.", &wg)

	fmt.Println("---------- Challenge Function ----------")

	// run the challenge function
	challenge()
}
