package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/steffantucker/holdsAPI/handlers/comics"
)

func main() {
	s, err := newserver()
	if err != nil {
		fmt.Println("encountered error", err)
	}
	ch := comics.NewComicsHandler(s.logger)

	getComicsRouter := s.router.Methods(http.MethodGet).Subrouter()
	getComicsRouter.HandleFunc("/comics", ch.GetAllComics)

	postComicsRouter := s.router.Methods(http.MethodPost).Subrouter()
	postComicsRouter.HandleFunc("/comics", ch.NewComic)
	postComicsRouter.Use(ch.MiddlewareVerifyComicData)

	putComicsRouter := s.router.Methods(http.MethodPut).Subrouter()
	putComicsRouter.HandleFunc("/comics/{id:[0-9]+}", ch.UpdateComic)
	putComicsRouter.Use(ch.MiddlewareVerifyComicData)

	s.start()

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
	s.shutdown(ctx)

}
