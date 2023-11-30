package models

// Artist : Main struct
type Artist struct {
	ID              uint     `json:"id"`
	Image           string   `json:"image,omitempty"`
	Name            string   `json:"name"`
	Members         []string `json:"members"`
	CreationDate    uint     `json:"creationDate"`
	FirstAlbum      string   `json:"firstAlbum"`
	LocationsURL    string   `json:"locations"`
	ConcertDatesURL string   `json:"concertDates"`
	RelationsURL    string   `json:"relations"`
}
