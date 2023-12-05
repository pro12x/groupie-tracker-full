package handlers

import (
	"encoding/json"
	"net/http"
)

func Search(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	Art, err := GetArtists()
	if err != nil {
		errorResponse(w, http.StatusInternalServerError)
		return
	}
	search := GetSearch(query, Art)
	if len(search.Artists) == 0 && len(search.Members) == 0 && len(search.Locations) == 0 && len(search.FirstAlbums) == 0 && len(search.CreationDates) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	jsonSearch, err := json.Marshal(search)
	if err != nil {
		errorResponse(w, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonSearch)
	//renderTemplates(w, "search", &models.MainData{AppInfos: models.App{AppName: appName, PageTitle: "Search", Attr: "search"}})
}
