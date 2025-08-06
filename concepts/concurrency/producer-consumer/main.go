package main

const NumberOfPizzas = 10

var pizzasMade, pizzasFailed, pizzasTotal int

type Producer struct {
	data chan PizzaOrder
	quit chan chan error
}

type PizzaOrder struct {
	pizzaNumber int
	message     string
	success     bool
}

func main() {
	// seed the random number generator

	// print out a message

	// create the producer

	// run the producer in the background

	// create and run the consumer

	// print out the ending message
}
