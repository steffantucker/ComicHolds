package series

import "net/http"

// GetAllSeries handles GET requests for all series
// shall be paginated
func GetAllSeries(w http.ResponseWriter, r *http.Request) {}

// NewSeries adds a new series to the database
func NewSeries(w http.ResponseWriter, r *http.Request) {}

// GetSeries handles GET requests for a specific series
func GetSeries(w http.ResponseWriter, r *http.Request) {}
