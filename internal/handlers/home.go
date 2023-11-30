package handlers

import (
	"bim/internal/models"
	"net/http"
)

// Home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		if r.URL.Path != "/" {
			errorResponse(w, http.StatusNotFound)
			return
		}
		renderTemplates(w, "home", &models.MainData{AppInfos: models.App{AppName: appName, PageTitle: "Home", Attr: "home"}})
	} else {
		errorResponse(w, http.StatusMethodNotAllowed)
	}
}
