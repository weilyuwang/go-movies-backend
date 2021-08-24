package api

import (
	"backend/models"
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	// To import a package solely for its side-effects (initialization),
	// use the blank identifier as explicit package name.
	// the underscore import is used for the side-effect of registering the postgresql driver
	// as a database driver in the init() function, without importing any other functions:
	_ "github.com/lib/pq"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
	db   struct {
		dsn string
	}
	jwt struct {
		secret string
	}
}

// In Go, an identifier that starts with a capital letter is exported from the package
type AppStatus struct {
	Status      string `json:"status"`
	Environment string `json:"environment"`
	Version     string `json:"version"`
}

type application struct {
	config config
	logger *log.Logger
	models models.Models
}

func RunServer() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "Server port to listen on")
	flag.StringVar(&cfg.env, "env", "development", "Applicaton environment (development|production)")
	// If use locally-installed postgres server, set host to localhost
	flag.StringVar(&cfg.db.dsn, "dsn", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"postgres", 5432, "postgres", "postgres", "go_movies"), "postgres connection string")
	flag.Parse()

	// go run main.go -h                                                                                      ─╯
	// -dsn string
	//       postgres connection string (default "postgres://postgres:postgres@localhost/go_movies?sslmode=disable")
	// -env string
	//       Applicaton environment (development|production) (default "development")
	// -port int
	//       Server port to listen on (default 4000)

	// read JWT secret from env
	cfg.jwt.secret = os.Getenv("GO_MOVIES_JWT")

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	if cfg.jwt.secret == "" {
		logger.Fatal("No JWT Secret specified")
	}

	db, err := openDB(cfg)
	if err != nil {
		logger.Fatal(err)
	}
	logger.Println("Connected to DB")
	defer db.Close()

	app := &application{
		config: cfg,
		logger: logger,
		models: models.NewModels(db),
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Println("Starting server on port", cfg.port)

	err = srv.ListenAndServe()
	if err != nil {
		log.Println(err)
	}
}

func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
