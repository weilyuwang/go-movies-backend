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
(sudo -u postgres) psql -d go_movies -f db_migration/go_movies.sql
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
pg_dump --no-owner go_movies > db_migration/go_movies.sql
```

### Docker

Rebuild and start: `docker-compose up --build`

Login: `docker exec -it postgres psql -U postgres`

Show tables: `\dt`

Show databases: `\l`

backup database: `docker exec -t your-db-container pg_dumpall -c -U postgres > dump.sql`

restore database from sql dump: `cat your_dump.sql | docker exec -i your-db-container psql -U postgres`

e.g. `cat db_migration/go_movies.sql | docker exec -i postgres psql -U postgres -d go_movies` (use -d to specify DB name)
