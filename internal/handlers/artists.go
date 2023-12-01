package handlers

import (
	"bim/internal/models"
	"bim/internal/pkg"
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
		for i := range all {
			all[i].FirstAlbum = pkg.FormatDate(all[i].FirstAlbum)
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
	// Get the relation
	relation, err2 := GetRelation(ID)
	// Get the location
	location, err3 := GetLocation(ID)
	// Get the date
	date, err4 := GetDate(ID)
	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		errorResponse(w, http.StatusInternalServerError)
		return
	}
	date = pkg.ProcessDates(date)
	location = pkg.ProcessString(location)
	artist.FirstAlbum = pkg.FormatDate(artist.FirstAlbum)
	relation = pkg.ProcessDatesLocations(relation)
	renderTemplates(w, "infos", &models.MainData{AppInfos: models.App{AppName: appName, PageTitle: strconv.Itoa(ID) + " - " + artist.Name, Attr: "artists"}, ArtistOne: artist, RelationOne: relation, LocationOne: location, DateOne: date})
}
