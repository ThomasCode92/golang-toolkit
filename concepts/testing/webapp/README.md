# The `webapp` Module - Web Application & REST API

`webapp` is a single Go module containing a server-rendered web application, a
JSON REST API, and a small helper CLI. Both servers use PostgreSQL and share the
same data and repository layers.

## 📁 Project Layout

```bash
webapp/
├── cmd/
│   ├── api/         # JSON REST API server (port 8090)
│   ├── web/         # server-rendered web app (port 8080)
│   └── cli/         # generates JWTs for manual API testing
├── pkg/
│   ├── data/        # domain models
│   └── repository/  # repository interface + Postgres/test implementations
├── templates/       # HTML templates for the web app
├── sql/init.sql     # database schema and seeded admin user
├── api/api.http     # sample API requests for an IDE HTTP client
└── compose.yml      # PostgreSQL and live-reloading webapp container
```

## 🧩 Components

| Path      | Port | Description                                                |
|-----------|------|------------------------------------------------------------|
| `cmd/web` | 8080 | Server-rendered web app using `html/template` and sessions |
| `cmd/api` | 8090 | JSON REST API using JWT access and refresh tokens          |
| `cmd/cli` | —    | Prints signed JWTs for manual API testing                  |

The repository pattern keeps handlers testable:

- `pkg/repository/repository.go` defines the `DatabaseRepo` interface.
- `pkg/repository/dbrepo/users_postgres.go` is the production implementation.
- `pkg/repository/dbrepo/users_testdb.go` is a fake used by unit tests.

## 🌐 Web Application (`cmd/web`)

| Method | Path                           | Description                      |
|--------|--------------------------------|----------------------------------|
| `GET`  | `/`                            | Home page with a login form      |
| `POST` | `/login`                       | Authenticate and start a session |
| `GET`  | `/user/profile`                | Protected profile page           |
| `POST` | `/user/upload-profile-picture` | Protected upload, 5 MB maximum   |
| `GET`  | `/static/*`                    | Static assets                    |

The database is seeded with an admin user:

- Email: `admin@example.com`
- Password: `secret`

## 🔐 REST API (`cmd/api`)

| Method   | Path              | Authentication | Description                                |
|----------|-------------------|----------------|--------------------------------------------|
| `POST`   | `/auth`           | —              | JSON login; returns an access/refresh pair |
| `POST`   | `/refresh-token`  | —              | Form `refresh_token`; returns a new pair   |
| `GET`    | `/users`          | Bearer token   | List all users                             |
| `GET`    | `/users/{userID}` | Bearer token   | Get one user                               |
| `PUT`    | `/users`          | Bearer token   | Insert a user                              |
| `PATCH`  | `/users/{userID}` | Bearer token   | Update a user                              |
| `DELETE` | `/users/{userID}` | Bearer token   | Delete a user                              |

Access tokens expire after 15 minutes and refresh tokens after 24 hours. A
refresh request is only accepted when the refresh token is within 30 seconds of
expiring. Protected endpoints expect the header:

```http
Authorization: Bearer <access_token>
```

The `cmd/cli` helper uses the same default signing secret and issuer (`example.com`) as the API, so its tokens work with
the default API settings.

## 🚀 Running the Applications

### Prerequisites

- [Go 1.24+](https://go.dev/dl/)
- [Docker](https://www.docker.com/get-started) with Docker Compose

### Option 1: Everything in Docker (recommended)

```bash
cd concepts/testing/webapp

docker compose up -d
docker compose logs -f webapp  # follow build and server logs
```

Compose starts PostgreSQL and an `air` container. Air builds and runs both
servers and rebuilds them when source files change. The applications are then
available at:

- Web app: <http://localhost:8080>
- REST API: <http://localhost:8090>

Stop the containers with:

```bash
docker compose down

# also remove the PostgreSQL data volume (re-runs sql/init.sql next time)
docker compose down -v
```

If you change `.air.toml`, recreate the container so Air picks up the new
configuration:

```bash
docker compose up -d --force-recreate webapp
```

### Option 2: Run the Servers on Your Host

Start only PostgreSQL in Docker, then run each server in its own terminal. The
default DSN uses the hostname `postgres`, which only resolves inside the Compose
network, so pass `-dsn` pointing to `localhost` instead:

```bash
cd concepts/testing/webapp
docker compose up -d postgres

DSN="host=localhost port=5432 user=postgres password=postgres dbname=users sslmode=disable timezone=UTC connect_timeout=5"

# terminal 1 - web app on :8080
go run ./cmd/web -dsn "$DSN"

# terminal 2 - REST API on :8090
go run ./cmd/api -dsn "$DSN"
```

The API also accepts `-domain` (used as the JWT issuer/audience) and
`-jwt_secret` flags; both have sensible defaults for local development.

## 🧪 Manual API Testing

`api/api.http` contains sample requests for IDE HTTP clients such as the
JetBrains HTTP client or VS Code REST Client. You can also use `curl`.

From `concepts/testing/webapp`, generate a valid token with the CLI and call a
protected endpoint:

```bash
TOKEN=$(go run ./cmd/cli -action=valid | tail -n 1)

curl -s http://localhost:8090/users \
  -H "Authorization: Bearer $TOKEN"
```

Use `-action=expired` to generate an expired token for testing failure cases.
Authenticate with the seeded admin user:

```bash
curl -s -X POST http://localhost:8090/auth \
  -H "Content-Type: application/json" \
  -d '{"email":"admin@example.com","password":"secret"}'
```

## ✅ Running the Tests

The handler and unit tests use the in-memory `TestDBRepo`, so they do not
require Docker:

```bash
cd concepts/testing/webapp

go test ./...              # run all unit and handler tests
go test -cover ./...       # show coverage per package

go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

Integration tests run against a real PostgreSQL instance started with
`dockertest`:

```bash
go test -tags=integration ./...
```

`dockertest` connects through the `DOCKER_HOST` environment variable; it does
not use the context selected by the `docker` CLI. If your CLI uses Docker
Desktop's `desktop-linux` context while `DOCKER_HOST` is unset, the tests try
the default `/var/run/docker.sock` and fail with:

```text
dial unix /var/run/docker.sock: connect: no such file or directory
```

Point `DOCKER_HOST` at the Docker Desktop socket before running the tests:

```bash
export DOCKER_HOST="unix://$HOME/.docker/desktop/docker.sock"
go test -tags=integration ./...
```

In fish:

```fish
set -gx DOCKER_HOST "unix://$HOME/.docker/desktop/docker.sock"
go test -tags=integration ./...
```

> The integration tests bind to host port `5432`. Stop the Compose database
> first (`docker compose down`) or the tests will not be able to start their own
> PostgreSQL container.
