### Start the backend server

```
export GO_MOVIES_JWT='[YOUR_JWT_SECRET]'
go run ./cmd/api
```

### Postgres DB

#### Start DB before running the server

- Before running the server, make sure the postgres is running

#### SQL dump

- create a database named `go_movies`, then in the project root directory, run:

```
psql -d go_movies -f sql_dump/go_movies.sql
```

### Production Build

- To generate a single binary executable go application:

```
env GOOS=linux GOARCH=amd64 go build -o gomovies ./cmd/api
```

### Dump the existing database

```
pg_dump  --no-owner go_movies > go_movies.sql
```
