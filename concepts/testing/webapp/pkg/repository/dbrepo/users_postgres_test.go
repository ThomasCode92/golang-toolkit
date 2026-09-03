package dbrepo

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"
	"time"
	"webapp/pkg/data"
	"webapp/pkg/repository"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = "postgres"
	dbname   = "users_test"
	port     = "5432"
	dsn      = "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable timezone=UTC connect_timeout=5"
)

var resource *dockertest.Resource
var pool *dockertest.Pool
var testDB *sql.DB
var testRepo repository.DatabaseRepo

func TestMain(m *testing.M) {
	// connect to docker; fail if not running
	p, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("could not connect to docker; is it running? %s", err)
	}

	pool = p

	// set up docker options, specify image and so forth
	opts := dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "14.5",
		Env: []string{
			"POSTGRES_USER=" + user,
			"POSTGRES_PASSWORD=" + password,
			"POSTGRES_DB=" + dbname,
		},
		ExposedPorts: []string{"5432"},
		PortBindings: map[docker.Port][]docker.PortBinding{
			"5432": {
				{HostIP: "0.0.0.0", HostPort: "5432"},
			},
		},
	}

	// get a resource (docker image)
	resource, err = pool.RunWithOptions(&opts)
	if err != nil {
		_ = pool.Purge(resource)
		log.Fatalf("could not start resource: %s", err)
	}

	// start the container, wait until ready
	if err := pool.Retry(connect); err != nil {
		_ = pool.Purge(resource)
		log.Fatalf("could not connect to database: %s", err)
	}

	// populate the database with empty tables
	err = createTables()
	if err != nil {
		log.Fatalf("error creating tables: %s", err)
	}

	testRepo = &PostgresDBRepo{
		DB: testDB,
	}

	// run tests
	code := m.Run()

	// clean up
	if err := pool.Purge(resource); err != nil {
		log.Fatalf("could not purge resource: %s", err)
	}

	os.Exit(code)
}

func connect() error {
	var err error
	testDB, err = sql.Open("pgx", fmt.Sprintf(dsn, host, port, user, password, dbname))
	if err != nil {
		log.Println("Error:", err)
		return err
	}
	return testDB.Ping()
}

func createTables() error {
	tableSQL, err := os.ReadFile("./testdata/users.sql")
	if err != nil {
		fmt.Println(err)
		return err
	}

	_, err = testDB.Exec(string(tableSQL))
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func Test_pingDB(t *testing.T) {
	err := testDB.Ping()
	if err != nil {
		t.Error("could not ping database")
	}
}

func TestPostgresDBRepo_InsertUser(t *testing.T) {
	testUser := data.User{
		Username:  "Admin",
		Email:     "admin@example.com",
		IsAdmin:   true,
		Password:  "secret",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	id, err := testRepo.InsertUser(testUser)
	if err != nil {
		t.Errorf("insert user returned an error: %s", err)
	}

	if id != 1 {
		t.Errorf("insert user returned wrong id; expected 1, got %d", id)
	}
}

func TestPostgresDBRepo_AllUsers(t *testing.T) {
	users, err := testRepo.AllUsers()
	if err != nil {
		t.Errorf("all users returned an error: %s", err)
	}

	if len(users) != 1 {
		t.Errorf("all users returned wrong number of users; expected 1, got %d", len(users))
	}

	testUser := data.User{
		Username:  "Jack",
		Email:     "jack@example.com",
		IsAdmin:   false,
		Password:  "secret",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	testRepo.InsertUser(testUser)

	users, err = testRepo.AllUsers()
	if err != nil {
		t.Errorf("all users returned an error: %s", err)
	}

	if len(users) != 2 {
		t.Errorf("all users returned wrong number of users after insert; expected 2, got %d", len(users))
	}
}

func TestPostgresDBRepo_GetUser(t *testing.T) {
	user, err := testRepo.GetUser(1)
	if err != nil {
		t.Errorf("get user returned an error: %s", err)
	}

	if user.Email != "admin@example.com" {
		t.Errorf("get user returned wrong user; expected admin@example.com, got %s", user.Email)
	}

	_, err = testRepo.GetUser(3)
	if err == nil {
		t.Error("no error returned when getting non-existent user")
	}
}

func TestPostgresDBRepo_GetUserByEmail(t *testing.T) {
	user, err := testRepo.GetUserByEmail("jack@example.com")
	if err != nil {
		t.Errorf("get user by email returned an error: %s", err)
	}

	if user.Username != "Jack" {
		t.Errorf("get user by email returned wrong user; expected Jack, got %s", user.Username)
	}
}

func TestPostgresDBRepo_UpdateUser(t *testing.T) {
	user, _ := testRepo.GetUser(2)
	user.Username = "Jane"
	user.Email = "jane@example.com"

	err := testRepo.UpdateUser(*user)
	if err != nil {
		t.Errorf("update user returned an error: %s", err)
	}

	updatedUser, _ := testRepo.GetUser(2)
	if updatedUser.Username != "Jane" || updatedUser.Email != "jane@example.com" {
		t.Errorf("update user did not correctly update the user; got %+v", updatedUser)
	}
}

func TestPostgresDBRepo_DeleteUser(t *testing.T) {
	err := testRepo.DeleteUser(2)
	if err != nil {
		t.Errorf("delete user returned an error: %s", err)
	}

	_, err = testRepo.GetUser(2)
	if err == nil {
		t.Error("no error returned when getting deleted user")
	}
}
