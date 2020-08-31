package comics

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/steffantucker/holdsAPI/data"
)

// ComicHandler type
type ComicHandler struct {
	l *log.Logger
}

// NewComicsHandler attaches the logger to all requests
func NewComicsHandler(l *log.Logger) *ComicHandler {
	return &ComicHandler{l}
}

// GetAllComics returns all comics in the database
// let middleware handle filters and pages
func (ch *ComicHandler) GetAllComics(w http.ResponseWriter, r *http.Request) {
	// need to query database for all comics
	// limit by count through middleware at some point
	ch.l.Println("Handle GET comics")
	d := data.GetComics()
	err := d.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to marshall JSON data", http.StatusInternalServerError)
	}
}

// NewComic will insert a new comic into the database
func (ch *ComicHandler) NewComic(w http.ResponseWriter, r *http.Request) {
	ch.l.Println("Handle POST comics")
	comic := r.Context().Value(KeyComic{}).(data.Comic)
	data.AddComic(&comic)
}

// UpdateComic updates the data associated with a comic
func (ch *ComicHandler) UpdateComic(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Unable to convert ID", http.StatusBadRequest)
	}

	ch.l.Println("Handle PUT comics", id)
	data := r.Context().Value(KeyComic{}).(data.Comic)

	err = data.UpdateProduct(id, &data)
	if err == data.ComicNotFound {
		http.Error(w, "Comic not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "Comic not found", http.StatusInternalServerError)
	}
}

// GetComic will return 1 comic based on ID
// because of middleware, ID is stored in context
func (ch *ComicHandler) GetComic(w http.ResponseWriter, r *http.Request) {
	// need to get the ID from context, then get
	// comic from DB by ID
}

// KeyComic is
type KeyComic struct{}

// MiddlewareVerifyComicData is middleware to verify incoming data
func (ch ComicHandler) MiddlewareVerifyComicData(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		comic := data.Comic{}
		err := comic.FromJSON(r.Body)
		if err != nil {
			ch.l.Println("[ERROR] deserializing product", err)
			http.Error(w, "Error reading comic data", http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), KeyComic{}, comic)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
