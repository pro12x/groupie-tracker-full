package handlers

import (
	"bim/internal/models"
	"net/http"
)

func Relations(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		if r.URL.Path != "/relations" {
			errorResponse(w, http.StatusNotFound)
			return
		}
		all, err := GetRelations()
		if err != nil {
			errorResponse(w, http.StatusInternalServerError)
			return
		}
		renderTemplates(w, "relations", &models.MainData{AppInfos: models.App{AppName: appName, PageTitle: "List of relations", Attr: "relations"}, Relations: all})
	}
}
