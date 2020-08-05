package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/steffantucker/holdsapi/postgres"
)

func initializeAPI() (*chi.Mux, *postgres.Db) {
	router := chi.NewRouter()

	db, err := postgres.New(
		postgres.ConnString("localhost", 5432, "comicholds", "password", "comicholds"),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// middleware for router
	router.Use(
		middleware.RequestID,
		middleware.Logger,
		middleware.Recoverer,
		middleware.URLFormat,
		render.SetContentType(render.ContentTypeJSON),
	)

	// Comic book routes
	router.Route("/comics", func(r chi.Router) {
		r.Get("/", comics.GetAllComics)
		r.Post("/", comics.NewComic)
		r.Get("/{id}", comics.GetComic)
	})
	router.Route("/series", func(r chi.Router) {
		r.With(paginate).Get("/", series.GetAllSeries)
		r.Post("/", series.NewSeries)
		r.Get("/{id}", series.GetSeries)
	})

	return router, db
}

func main() {
	router, db := initializeAPI()
	defer db.Close()

	log.Fatal(http.ListenAndServe(":8080", router))
}
