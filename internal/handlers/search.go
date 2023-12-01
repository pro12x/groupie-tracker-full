package handlers

import (
	"bim/internal/models"
	"net/http"
)

func Search(w http.ResponseWriter, r *http.Request) {
	renderTemplates(w, "search", &models.MainData{})
}
