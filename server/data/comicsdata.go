package data

import (
	"encoding/json"
	"fmt"
	"io"

	uuid "github.com/satori/go.uuid"
)

// Comic struct for comic book info
type Comic struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Issue       int       `json:"issue"`
	TotalIssues int       `json:"total_issues"`
	SeriesID    int       `json:"series_id"`
	PublisherID int       `json:"publisher_id"`
}

// ErrComicNotFound is the error if the requested comic isn't found
var ErrComicNotFound = fmt.Errorf("Comic not found")

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
func AddComic(newComic *Comic) error {
	// TODO heck to make sure there's no conflict with existing
	// comic. will need to check against name and also
	// against issue number in series
	newComic.ID = uuid.NewV4()
	comicsList = append(comicsList, newComic)
	return nil
}

// UpdateComic updates the requested comic
func UpdateComic(uid uuid.UUID, newData Comic) error {
	id := getIndexfromProductID(uid)
	if id == -1 {
		return ErrComicNotFound
	}
	comicsList[id] = &newData
	return nil
}

func getIndexfromProductID(id uuid.UUID) int {
	for i, c := range comicsList {
		if c.ID == id {
			return i
		}
	}
	return -1
}

var comicsList = Comics{
	&Comic{
		ID:          uuid.NewV4(),
		Name:        "Doom Patrol",
		Issue:       1,
		TotalIssues: 1,
		SeriesID:    1,
		PublisherID: 1,
	},
}
