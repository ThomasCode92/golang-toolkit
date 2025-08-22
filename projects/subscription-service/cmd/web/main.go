package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/redisstore"
	"github.com/alexedwards/scs/v2"
	"github.com/gomodule/redigo/redis"
	_ "github.com/jackc/pgx/v5/stdlib" // import the pgx driver for database/sql
)

const webPort = "8080"

func main() {
	// connect to the database
	db := initDB()
	db.Ping()

	// create sessions
	session := initSession()

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

func initSession() *scs.SessionManager {
	redisPool := initRedis()
	session := scs.New()

	session.Store = redisstore.New(redisPool)
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = true

	return session
}

func initRedis() *redis.Pool {
	return &redis.Pool{
		MaxIdle: 10,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", os.Getenv("REDIS_URL"))
		},
	}
}
