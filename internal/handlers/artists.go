package handlers

import (
	"bim/internal/models"
	"net/http"
	"strconv"
)

func Artists(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		if r.URL.Path != "/artists" {
			errorResponse(w, http.StatusNotFound)
			return
		}
		all, err := GetArtists()
		if err != nil {
			errorResponse(w, http.StatusBadRequest)
			return
		}
		renderTemplates(w, "artists", &models.MainData{AppInfos: models.App{AppName: appName, PageTitle: "List of Artists", Attr: "artists"}, ArtistList: all})
	}
}

func Artist(w http.ResponseWriter, r *http.Request) {
	// Check if artist exists
	if len(r.URL.Path) <= len("/artist/") {
		errorResponse(w, http.StatusBadRequest)
		return
	}
	// Extract ID
	id := r.URL.Path[len("/artist/"):]
	ID, err0 := strconv.Atoi(id)
	if err0 != nil {
		errorResponse(w, http.StatusBadRequest)
		return
	}
	// Restricted ID
	if ID < 1 || ID > 52 {
		errorResponse(w, http.StatusNotFound)
		return
	}
	// Get the artist
	artist, err1 := GetArtist(ID)
	if err1 != nil {
		errorResponse(w, http.StatusInternalServerError)
		return
	}
	// Get the relation
	relation, err2 := GetRelation(ID)
	if err2 != nil {
		errorResponse(w, http.StatusInternalServerError)
		return
	}
	renderTemplates(w, "infos", &models.MainData{AppInfos: models.App{AppName: appName, PageTitle: strconv.Itoa(ID) + " - " + artist.Name, Attr: "artists"}, ArtistOne: artist, RelationOne: relation})
}
