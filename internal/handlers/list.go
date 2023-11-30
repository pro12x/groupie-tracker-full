package handlers

import (
	"bim/internal/models"
	"net/http"
)

func List(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		if r.URL.Path != "/list" {
			errorResponse(w, http.StatusNotFound)
			return
		}
		list1, err1 := GetRelations()
		// fmt.Println(list1)
		if err1 != nil {
			errorResponse(w, http.StatusBadRequest)
			return
		}
		list2, err2 := GetLocations()
		// fmt.Println(list2)
		if err2 != nil {
			errorResponse(w, http.StatusBadRequest)
			return
		}
		list3, err3 := GetDates()
		// fmt.Println(list3)
		if err3 != nil {
			errorResponse(w, http.StatusBadRequest)
			return
		}
		renderTemplates(w, "list", &models.MainData{AppInfos: models.App{AppName: appName, PageTitle: "List", Attr: "list"}, Relations: list1, Locations: list2, Dates: list3})
	}
}
