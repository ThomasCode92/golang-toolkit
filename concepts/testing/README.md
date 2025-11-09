# Introduction to Testing in Go

Learn how to write effective unit and integration tests in Go, for web
applications and REST APIs

---

## 📚 Testing in Go

- Testing is a critical part of the development process.
- Well tested code is more maintainable, more secure, and overall takes less
  time to write.
- Go has a rich set of tools for testing built right integration.
- We wont't use any third party testing packages at all (e.g. Testify).
- We'll write unit tests and integration tests.

## ❓ About this project

> This project is based on the
> [Introduction to Testing in Go](https://www.udemy.com/course/introduction-to-testing-in-go-golang/)
> course of Travis Sawler.

### 🧠 Topics & Features

- Prime number CLI
  - A simple command line program that takes user input.
  - The program will determine whether or not a given number is prime.
  - We'll write tests for that program.
- Simple Web Applications
  - We'll build a simple web application that allows users to authenticate
    (securely).
  - We'll write some middleware to check authentication, and to add some data to
    the context.
  - We'll also use the _Repository Pattern_ to make the code more testable.
  - Write tests for all functionality, including integration tests.
- REST API
  - We'll write a simple REST API that allows authentication using JWT tokens.
  - We'll write some middleware to check authentication by looking at the
    Authorization header for incoming requests.
  - We'll create endpoints for typical CRUD operations.
  - Write tests for all functionality.
- Update REST API to support SPAs
  - We'll update our REST API to include functionality that makes it useable
    with any Single Page Application (SPA).
  - We'll add functionality to store JWTs and Refresh tokens securely.
  - We'll create endpoints for typical SPA functionality.
  - Write tests for all functionality.

### 📁 Structure

```
testing/
├── primeapp/     # simple prime number CLI application
└──
```

### 🚀 Running the Example Applications

Each subfolder is a Go project. To run the example applications, you can use the
following commands:

```bash
# primeapp is an example, replace it with any other subfolder
cd primeapp && go run .        # run the application
cd primeapp && go run -race .  # run, and check for race conditions
cd primeapp && go test -race . # run the tests, and check for race conditions
```
