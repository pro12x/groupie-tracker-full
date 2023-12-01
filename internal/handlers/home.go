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
		rand, _ := GetRandom(4)
		renderTemplates(w, "home", &models.MainData{AppInfos: models.App{AppName: appName, PageTitle: "Home", Attr: "home"}, Random: rand})
		return
	} /*else {
		errorResponse(w, http.StatusMethodNotAllowed)
	}*/
}
