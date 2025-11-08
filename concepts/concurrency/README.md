# Working with Concurrency in Golang

Learn the advantages and pitfalls of **concurrent programming** with the Go
programming language

---

## 📚 Concurrency in Go

**Don't communicate by sharing memory, share memory by communicating.**

Don't over-engineer things by using shared memory and complicated, error-prone
synchronization primitives; instead, use message-passing between goroutines so
variables and data can be used in the appropriate sequence.

- Golden rule for concurrency: if you don't need it, don't use it.
- Keep your application's complexity to an absolute minimum; it's easier to
  write, easier to understand, and easier to maintain.

## ❓ About this project

> This project is based on the
> [Working with Concurrency in Go](https://www.udemy.com/course/working-with-concurrency-in-go-golang)
> course of Travis Sawler.

### 🧠 Topics & Features

- Basic types of the _sync package_: mutexes, and wait groups
- Three classic computer science problems:
  - The _dining philosophers_ problem
  - The _producer-consumer_ problem
  - The _weekly income_ problem
- A more real-world scenario, a subset of a larger (subscription) problem.

### 📁 Structure

```
concurrency/
├── channels/           # an introduction to channels
├── dinning/            # an example of the dining philosophers problem
├── goroutine/          # an introduction to goroutines
├── mutex/              # an introduction to mutexes
├── producer-consumer/  # an example of the producer-consumer problem
├── sleeping-barber/    # an example of the sleeping barber problem
└── weekly-income/      # an example project, using goroutines and mutexes
```

### 🚀 Running the Example Applications

Each subfolder is a Go project. To run the example applications, you can use the
following commands:

```bash
# weekly-income is an example, replace it with any other subfolder
cd weekly-income && go run .        # run the application
cd weekly-income && go run -race .  # run, and check for race conditions
cd weekly-income && go test -race . # run the tests, and check for race conditions
```

**A more 'real-world' scenario:**<br />To see a more realistic example of
concurrency in Go, check out the
[subscription-service](/projects/subscription-service/README.md) project.
