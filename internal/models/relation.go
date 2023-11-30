package models

// Relation : One
type Relation struct {
	ID             uint                `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

// Relations : All
type Relations struct {
	Index []struct {
		ID             uint                `json:"id"`
		DatesLocations map[string][]string `json:"datesLocations"`
	} `json:"index"`
}
