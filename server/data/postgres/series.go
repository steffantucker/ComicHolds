package series

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi"
)

// SeriesHandler structure holds database connection
type SeriesHandler struct {
	Conn *sql.DB
}

// NewHandler creates new struct with db
func NewHandler(conn *sql.DB) *SeriesHandler {
	return &SeriesHandler{
		Conn: conn,
	}
}

// SeriesRouter handles chi routes for series
func (ch *SeriesHandler) SeriesRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/", GetAllSeries)
	r.Get("/{id:[0-9]+", GetSeries)
	r.Post("/", NewSeries)

	return r
}
