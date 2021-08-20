### Start the backend development server

```
export GO_MOVIES_JWT='[YOUR_JWT_SECRET]'
go run main.go
```

### Postgres DB

#### Start DB before running the server

- Before running the server, make sure the postgres is running

#### populate the databse with existing SQL dump file

- create a database named `go_movies`, then in the project root directory, run:

```
psql -d go_movies -f go_movies.sql

or

sudo -u postgres psql -d go_movies -f go_movies.sql
```

### Production Build

- To generate a single binary executable go application:

```
go build -o gomovies main.go
```

- For linux remote server:

```
env GOOS=linux GOARCH=amd64 go build -o gomovies main.go
```

### Dump the existing database

```
pg_dump  --no-owner go_movies > go_movies.sql
```
