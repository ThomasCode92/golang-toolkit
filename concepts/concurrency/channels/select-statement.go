package main

import (
	"fmt"
	"time"
)

func serverOne(ch chan string) {
	for {
		time.Sleep(6 * time.Second)
		ch <- "This is from server one"
	}
}

func serverTwo(ch chan string) {
	for {
		time.Sleep(3 * time.Second)
		ch <- "This is from server two"
	}
}

func servers() {
	fmt.Println("Select with Channels")
	fmt.Println("--------------------")

	channelOne := make(chan string)
	channelTwo := make(chan string)

	go serverOne(channelOne)
	go serverTwo(channelTwo)

	for {
		select {
		case s1 := <-channelOne:
			fmt.Println("Case one:", s1)
		case s2 := <-channelOne:
			fmt.Println("Case two", s2)
		case s3 := <-channelTwo:
			fmt.Println("Case three", s3)
		case s4 := <-channelTwo:
			fmt.Println("Case four", s4)
		default:
			// avoiding deadlock
		}
	}
}
