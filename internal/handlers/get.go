package handlers

import (
	"bim/internal/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	url2 "net/url"
)

const (
	arts = "https://groupietrackers.herokuapp.com/api/artists"
	art  = "https://groupietrackers.herokuapp.com/api/artists/"
	rels = "https://groupietrackers.herokuapp.com/api/relation"
	rel  = "https://groupietrackers.herokuapp.com/api/relation/"
	dats = "https://groupietrackers.herokuapp.com/api/dates"
	dat  = "https://groupietrackers.herokuapp.com/api/dates/"
	locs = "https://groupietrackers.herokuapp.com/api/locations"
	loc  = "https://groupietrackers.herokuapp.com/api/locations/"
)

// getJSON returns data from API
func getJSON(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(response.Body)
	output, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return output, nil
}

// GetArtist returns artist
func GetArtist(id int) (models.Artist, error) {
	var artist models.Artist
	idStr := fmt.Sprintf("%d", id)
	url := art + url2.PathEscape(idStr)
	body, err := getJSON(url)
	if err != nil {
		return artist, err
	}
	err = json.Unmarshal(body, &artist)
	if err != nil {
		return artist, err
	}
	return artist, nil
}

// GetArtists returns a list of artists
func GetArtists() ([]models.Artist, error) {
	url := arts
	body, err := getJSON(url)
	var artists []models.Artist
	if err != nil {
		return artists, err
	}
	err = json.Unmarshal(body, &artists)
	if err != nil {
		return artists, err
	}
	return artists, nil
}

// GetRelation returns relation
func GetRelation(id int) (models.Relation, error) {
	var relation models.Relation
	idStr := fmt.Sprintf("%d", id)
	url := rel + url2.PathEscape(idStr)
	body, err := getJSON(url)
	if err != nil {
		return relation, err
	}
	err = json.Unmarshal(body, &relation)
	if err != nil {
		return relation, err
	}
	return relation, nil
}

// GetRelations returns a list of relations
func GetRelations() (models.Relations, error) {
	url := rels
	body, err := getJSON(url)
	var relations models.Relations
	if err != nil {
		return relations, err
	}
	err = json.Unmarshal(body, &relations)
	if err != nil {
		return relations, err
	}
	return relations, nil
}

// GetLocation returns location
func GetLocation(id int) (models.Location, error) {
	var location models.Location
	idStr := fmt.Sprintf("%d", id)
	url := loc + url2.PathEscape(idStr)
	body, err := getJSON(url)
	if err != nil {
		return location, err
	}
	err = json.Unmarshal(body, &location)
	if err != nil {
		return location, err
	}
	return location, nil
}

// GetLocations returns a list of locations
func GetLocations() (models.Locations, error) {
	url := locs
	body, err := getJSON(url)
	var locations models.Locations
	if err != nil {
		return locations, err
	}
	err = json.Unmarshal(body, &locations)
	if err != nil {
		return locations, err
	}
	return locations, nil
}

// GetDate returns date
func GetDate(id int) (models.Date, error) {
	var date models.Date
	idStr := fmt.Sprintf("%d", id)
	url := dat + url2.PathEscape(idStr)
	body, err := getJSON(url)
	if err != nil {
		return date, err
	}
	err = json.Unmarshal(body, &date)
	if err != nil {
		return date, err
	}
	return date, nil
}

// GetDates returns a list of dates
func GetDates() (models.Dates, error) {
	url := dats
	body, err := getJSON(url)
	var dates models.Dates
	if err != nil {
		return dates, err
	}
	err = json.Unmarshal(body, &dates)
	if err != nil {
		return dates, err
	}
	return dates, nil
}
