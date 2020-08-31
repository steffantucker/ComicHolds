package comics

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi"
)

// ComicHandler structure holds database connection
type ComicHandler struct {
	Conn *sql.DB
}

// NewHandler creates new struct with db
func NewHandler(conn *sql.DB) *ComicHandler {
	return &ComicHandler{
		Conn: conn,
	}
}

// ComicRouter handles chi routes for series
func (ch *ComicHandler) ComicRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/", GetAllComics)
	r.Get("/{id:[0-9]+", GetComic)
	r.Post("/", NewComic)

	return r
}
