package models

// Dates : All
type Dates struct {
	Index []struct {
		ID    uint     `json:"id"`
		Dates []string `json:"dates"`
	} `json:"index"`
}

// Date : One
type Date struct {
	ID    uint     `json:"id"`
	Dates []string `json:"dates"`
}
