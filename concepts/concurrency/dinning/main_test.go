package main

import (
	"testing"
	"time"
)

func Test_dine(t *testing.T) {
	eatTime = 0 * time.Second   // set to 0 to speed up the output
	thinkTime = 0 * time.Second // set to 0 to speed up the output
	sleepTime = 0 * time.Second // set to 0 to speed up the output

	for range 10 {
		orderFinished = []string{} // reset the orderFinished slice
		dine()
		if len(orderFinished) != len(philosophers) {
			t.Errorf("incorrect length of slice: expected %d but got %d", len(philosophers), len(orderFinished))
		}
	}
}

func Test_dineWithVaryingDelays(t *testing.T) {
	theTests := []struct {
		name  string
		delay time.Duration
	}{
		{"zero delay", time.Second * 0},
		{"quarter second delay", time.Millisecond * 250},
		{"half second delay", time.Millisecond * 500},
	}

	for _, e := range theTests {
		orderFinished = []string{}

		eatTime = e.delay
		sleepTime = e.delay
		thinkTime = e.delay

		dine()
		if len(orderFinished) != len(philosophers) {
			t.Errorf("%s: incorrect length of slice; expected %d but got %d", e.name, len(philosophers), len(orderFinished))
		}
	}
}
