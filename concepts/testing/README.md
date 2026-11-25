# Introduction to Testing in Go

Learn how to write effective unit and integration tests in Go, for web
applications and REST APIs

---

## 📚 Testing in Go

- Testing is a critical part of the development process.
- Well tested code is more maintainable, more secure, and overall takes less
  time to write.
- Go has a rich set of tools for testing built right integration.
- We won't use any third party testing packages at all (e.g. Testify).
- We'll write unit tests and integration tests.

## ❓ About this project

> This project is based on the
> [Introduction to Testing in Go](https://www.udemy.com/course/introduction-to-testing-in-go-golang/)
> course of Travis Sawler.

### 🧠 Topics & Features

- Prime number CLI
    - A simple command line program that takes user input.
    - The program will determine whether a given number is prime.
    - We'll write tests for that program.
- Simple Web Applications
    - We'll build a simple web application that allows users to authenticate (securely).
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
    - We'll update our REST API to include functionality that makes it usable
      with any Single Page Application (SPA).
    - We'll add functionality to store JWTs and Refresh tokens securely.
    - We'll create endpoints for typical SPA functionality.
    - Write tests for all functionality.

### 📁 Structure

```bash
testing/
├── primeapp/   # prime number CLI
└── webapp/     # web app + REST API (single Go module)
```

See [webapp/README.md](webapp/README.md) for the webapp architecture, routes,
configuration, manual API testing, and detailed test instructions.

## 🚀 Quick Start

### Prime Number CLI

```bash
cd concepts/testing/primeapp
go run .
```

### Web App & REST API

```bash
cd concepts/testing/webapp
docker compose up -d
```

- Web app: <http://localhost:8080>
- REST API: <http://localhost:8090>

## 🧪 Tests

```bash
# Prime number CLI
cd concepts/testing/primeapp
go test ./...

# Web app and API
cd concepts/testing/webapp
go test ./...
go test -tags=integration ./...  # requires Docker and a free port 5432
```
