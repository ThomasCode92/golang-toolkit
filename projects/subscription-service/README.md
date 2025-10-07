# Subscription Service - Using Concurrency

This project demonstrates the use of concurrency in Go. It showcases best
practices for managing concurrent operations in a scalable and efficient web
application.

> This project is based on the
> [Working with Concurrency in Go](https://www.udemy.com/course/working-with-concurrency-in-go-golang)
> course of Travis Sawler.

## 🛠️ Prerequisites

- [Go Programming Language](https://golang.org/dl/)
- [Docker](https://www.docker.com/get-started)

## 🚀 **Getting Started**

Follow these steps to set up and run the project:

1. **Start the development environment using Docker:**

   ```bash
   docker compose up -d
   ```

2. **Set the appropriate environment variables:**

   ```bash
   export DSN="host=localhost port=5432 user=postgres password=password dbname=sub-service"
   export REDIS_URL="localhost:6379"

   # when using Nushell, use the following command instead:
   load-env { "DSN": "host=localhost port=5432 user=postgres password=password dbname=sub-service", "REDIS_URL": "localhost:6379" }

   ```

3. **Set up the database schema:**

   ```bash
   psql "host=localhost port=5432 user=postgres password=password dbname=sub-service" -f ./db.sql
   ```

4. **Start the web server:**

   ```bash
   go run ./cmd/web/*.go
   ```

5. **Authenticate the dummy user:**

   ```bash
   # open the browser on http://localhost:8080/login
   Username: admin@example.com
   Password: verysecret
   ```
