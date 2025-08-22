package main

import (
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib" // import the pgx driver for database/sql
)

const webPort = "8080"

func main() {
	// connect to the database
	db := initDB()
	db.Ping()

	// create sessions

	// create channels

	// create waitGroup

	// set up the application configuration

	// set up mail

	// listen for web connections
}

func initDB() *sql.DB {
	conn := connectToDB()
	if conn == nil {
		log.Panic("can't connect to the database")
	}

	return conn
}

func connectToDB() *sql.DB {
	counts := 0

	for {
		conn, err := openDB(os.Getenv("DSN"))
		if err != nil {
			log.Println("database not yet ready...")
		} else {
			log.Println("connected to the database")
			return conn
		}

		counts++
		if counts > 10 {
			return nil
		}

		time.Sleep(1 * time.Second)
		continue
	}
}

func openDB(dsn string) (*sql.DB, error) {
	conn, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	return conn, conn.Ping()
}
