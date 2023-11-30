package models

// Locations : All
type Locations struct {
	Index []struct {
		ID        uint     `json:"id"`
		Locations []string `json:"locations"`
		DatesURL  string   `json:"dates"`
	} `json:"index"`
}

// Location : One
type Location struct {
	ID        uint     `json:"id"`
	Locations []string `json:"locations"`
	DatesURL  string   `json:"dates"`
}
