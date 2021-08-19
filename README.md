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
