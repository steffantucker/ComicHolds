package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/steffantucker/holdsAPI/comics"
	"github.com/steffantucker/holdsAPI/series"
)

func initializeAPI() (*chi.Mux, *sql.DB) {
	router := chi.NewRouter()

	db, err := initializeDB()
	if err != nil {
		log.Fatal(err)
	}

	// middleware for router
	router.Use(
		middleware.RequestID,
		middleware.Logger,
		middleware.Recoverer,
		render.SetContentType(render.ContentTypeJSON),
	)

	// Comic book routes
	comicHandler := comics.NewHandler(db)
	seriesHandler := series.NewHandler(db)
	router.Route("/", func(r chi.Router) {
		r.Mount("/comics", comicHandler.ComicRouter())
		r.Mount("/series", seriesHandler.SeriesRouter())
	})

	router.Route("/series", func(r chi.Router) {
		r.Get("/", series.GetAllSeries)
		r.Post("/", series.NewSeries)
		r.Get("/{id}", series.GetSeries)
	})

	return router, db
}

func initializeDB() (*sql.DB, error) {
	// TODO: move login info to env variables
	host := "localhost"
	port := 5432
	user := "comicholds"
	pass := "password"
	dbname := "comicholds"
	connstring := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, pass, dbname)
	return sql.Open("postgres", connstring)
}

func main() {
	router, db := initializeAPI()
	defer db.Close()
	log.Fatal(http.ListenAndServe(":8080", router))
}
