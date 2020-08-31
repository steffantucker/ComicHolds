package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/steffantucker/holdsAPI/handlers/comics"
)

func main() {
	bindaddress := os.Getenv("BIND_ADDRESS")
	if bindaddress == "" {
		bindaddress = ":9090"
	}
	l := log.New(os.Stdout, "inventory-api", log.LstdFlags)
	ch := comics.NewComicsHandler(l)

	sm := mux.NewRouter()

	getComicsRouter := sm.Methods(http.MethodGet).Subrouter()
	getComicsRouter.HandleFunc("/comics", ch.GetAllComics)

	postComicsRouter := sm.Methods(http.MethodPost).Subrouter()
	postComicsRouter.HandleFunc("/comics", ch.NewComic)
	postComicsRouter.Use(ch.MiddlewareVerifyComicData)

	putComicsRouter := sm.Methods(http.MethodPut).Subrouter()
	putComicsRouter.HandleFunc("/comics/{id:[0-9]+}", ch.UpdateComic)
	putComicsRouter.Use(ch.MiddlewareVerifyComicData)

	// create a new server
	s := http.Server{
		Addr:         bindaddress,       // configure the bind address
		Handler:      sm,                // set the default handler
		ErrorLog:     l,                 // set the logger for the server
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}

	// start the server
	go func() {
		l.Println("Starting server on port 9090")

		err := s.ListenAndServe()
		if err != nil {
			l.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block until a signal is received.
	sig := <-c
	log.Println("Got signal:", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	log.Println("Shutting down")
	s.Shutdown(ctx)

}
