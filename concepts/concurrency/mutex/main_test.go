package main

import (
	"sync"
	"testing"
)

func Test_updateMessage(t *testing.T) {
	var mutex sync.Mutex

	msg = "Hello, World!"

	wg.Add(2)
	go updateMessage("Goodbye, World!", &mutex)
	go updateMessage("Hello, Universe!", &mutex)
	wg.Wait()

	if msg != "Goodbye, World!" && msg != "Hello, Universe!" {
		t.Error("incorrect value in msg")
	}
}
