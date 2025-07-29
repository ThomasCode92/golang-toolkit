package main

import (
	"fmt"
	"time"
)

func printSomething(s string) {
	fmt.Println(s)
}

func main() {
	go printSomething("This is the first thing to be printed.")

	time.Sleep(1 * time.Second) // Wait for the goroutine to finish

	printSomething("This is the second thing to be printed.")
}
