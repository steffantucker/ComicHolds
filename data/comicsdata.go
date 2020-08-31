package data

import (
	"encoding/json"
	"io"
)

// Comic struct for comic book info
type Comic struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Issue       int    `json:"issue"`
	TotalIssues int    `json:"total_issues"`
	SeriesID    int    `json:"series_id"`
	PublisherID int    `json:"publisher_id"`
}

// Comics slice of multiple comics
type Comics []*Comic

// ToJSON serializes data to JSON
func (c *Comics) ToJSON(w io.Writer) error {
	return json.NewEncoder(w).Encode(c)
}

// FromJSON deserializes data from JSON
func (c *Comic) FromJSON(r io.Reader) error {
	return json.NewDecoder(r).Decode(c)
}

// GetComics returns all comics
func GetComics() Comics {
	return comicsList
}

// AddComic adds a comic to the database
func AddComic(newComic *Comic) {
	// TODO heck to make sure there's no conflict with existing
	// comic. will need to check against name and also
	// against issue number in series
	newComic.ID = getNextID()
	comicsList = append(comicsList, newComic)
}

func getNextID() int {
	lastID := comicsList[len(comicsList)-1].ID
	return lastID + 1
}

var comicsList = Comics{
	&Comic{
		ID:          1,
		Name:        "Doom Patrol",
		Issue:       1,
		TotalIssues: 1,
		SeriesID:    1,
		PublisherID: 1,
	},
}
