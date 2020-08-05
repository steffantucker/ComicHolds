package comics

import "net/http"

// Comic struct for comic book info
type Comic struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Issue       int    `json:"issue"`
	TotalIssues int    `json:"total_issues"`
	SeriesID    int    `json:"series_id"`
	Series      string `json:"series"`
	PublisherID int    `json:"publisher_id"`
	Publisher   string `json:"publisher"`
}

// GetAllComics returns all comics in the database
// let middleware handle filters and pages
func GetAllComics(w http.ResponseWriter, r *http.Request) {
	// need to query database for all comics
	// limit by count through middleware at some point

}

// NewComic will insert a new comic into the database
func NewComic(w http.ResponseWriter, r *http.Request) {
	// check to make sure there's no conflict with existing
	// comic. will need to check against name and also
	// against issue number in series

}

// GetComic will return 1 comic based on ID
// because of middleware, ID is stored in context
func GetComic(w http.ResponseWriter, r *http.Request) {
	// need to get the ID from context, then get
	// comic from DB by ID
}
