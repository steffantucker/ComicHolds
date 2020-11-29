package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

type server struct {
	//db     *sqlx.DB
	server *http.Server
	logger *log.Logger
	router *mux.Router
}

func newserver() (server, error) {
	var s server
	s.logger = log.New(os.Stdout, "inventory-api", log.LstdFlags)
	s.router = mux.NewRouter()

	bindaddress := os.Getenv("BIND_ADDRESS")
	if bindaddress == "" {
		bindaddress = ":9090"
	}

	s.server = &http.Server{
		Addr:         bindaddress,       // configure the bind address
		Handler:      s.router,          // set the default handler
		ErrorLog:     s.logger,          // set the logger for the server
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}

	// TODO: env variable for db name
	// var err error
	// s.db, err = sqlx.Open("sqlite3", "comics.db")
	// if err != nil {
	// 	s.logger.Println("cannot connect to database", err)
	// 	return server{}, err
	// }
	// row := s.db.QueryRow("SELECT name FROM sqlite_master WHERE type='table'")
	// var tablename string
	// err = row.Scan(&tablename)
	// if err == sql.ErrNoRows {
	// 	s.createTables()
	// } else if err != nil {
	// 	s.logger.Println("database exist but ran into an error", err)
	// 	return server{}, err
	// }
	return s, nil
}

func (s *server) start() {
	go func() {
		// TODO: print env variable for bind address
		s.logger.Println("Starting server on port 9090")

		err := s.server.ListenAndServe()
		if err != nil {
			s.logger.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()
}

func (s *server) shutdown(ctx context.Context) {
	s.server.Shutdown(ctx)
}

func (s *server) createTables() {
	return
}
