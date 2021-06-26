### Start the application

```
go run cmd/api/*.go

or

go run ./cmd/api
```

### Postgres with Docker

- start postgresql
- create a database named `go_movies`, then in the project root directory, run:

```
psql -d go_movies -f sql_dump/go_movies.sql
```
