package main

import (
	"bufio"
	"io"
	"os"
	"strings"
	"testing"
)

func Test_intro(t *testing.T) {
	oldOut := os.Stdout  // save a copy of os.Stdout
	r, w, _ := os.Pipe() // create a read and write pipe
	os.Stdout = w        // set os.Stdout to the write pipe

	intro() // call the function that prints to stdout

	_ = w.Close()           // close the write pipe
	os.Stdout = oldOut      // reset os.Stdout to its original value
	out, _ := io.ReadAll(r) // read the output from the read pipe

	// perform the test
	if !strings.Contains(string(out), "Enter a whole number") {
		t.Errorf("intro text not correct, got %s", string(out))
	}
}

func Test_prompt(t *testing.T) {
	oldOut := os.Stdout  // save a copy of os.Stdout
	r, w, _ := os.Pipe() // create a read and write pipe
	os.Stdout = w        // set os.Stdout to the write pipe

	prompt() // call the function that prints to stdout

	_ = w.Close()           // close the write pipe
	os.Stdout = oldOut      // reset os.Stdout to its original value
	out, _ := io.ReadAll(r) // read the output from the read pipe

	// perform the test
	if string(out) != "-> " {
		t.Errorf("incorrect prompt: expected '->', but got %s", string(out))
	}
}

func Test_checkNumbers(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "empty", input: "", expected: "Please enter a whole number!"},
		{name: "zero", input: "0", expected: "0 is not prime, by definition!"},
		{name: "one", input: "1", expected: "1 is not prime, by definition!"},
		{name: "two", input: "2", expected: "2 is a prime number!"},
		{name: "three", input: "3", expected: "3 is a prime number!"},
		{name: "negative", input: "-1", expected: "Negative numbers are not prime, by definition!"},
		{name: "typed", input: "three", expected: "Please enter a whole number!"},
		{name: "decimal", input: "1.1", expected: "Please enter a whole number!"},
		{name: "quit", input: "q", expected: ""},
		{name: "QUIT", input: "Q", expected: ""},
		{name: "greek", input: "επτά", expected: "Please enter a whole number!"},
	}

	for _, e := range tests {
		input := strings.NewReader(e.input)
		reader := bufio.NewScanner(input)
		res, _ := checkNumbers(reader)

		if !strings.EqualFold(res, e.expected) {
			t.Errorf("%s: expected %s, but got %s", e.name, e.expected, res)
		}
	}
}

func Test_isPrime(t *testing.T) {
	primeTests := []struct {
		name     string
		testNum  int
		expected bool
		msg      string
	}{
		{"prime", 7, true, "7 is a prime number!"},
		{"not prime", 8, false, "8 is not a prime number because it is divisible by 2!"},
		{"zero", 0, false, "0 is not prime, by definition!"},
		{"one", 1, false, "1 is not prime, by definition!"},
		{"negative number", -11, false, "Negative numbers are not prime, by definition!"},
	}

	for _, e := range primeTests {
		result, msg := isPrime(e.testNum)
		if e.expected && !result {
			t.Errorf("%s: expected true but got false", e.name)
		}

		if !e.expected && result {
			t.Errorf("%s: expected false but got true", e.name)
		}

		if e.msg != msg {
			t.Errorf("%s: expected %s but got %s", e.name, e.msg, msg)
		}
	}
}
