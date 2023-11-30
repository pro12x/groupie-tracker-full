package handlers

import (
	"bim/config"
	"bim/internal/models"
	"bim/internal/pkg"
	"bytes"
	"net/http"
	"path/filepath"
	"text/template"
)

const appName = "Groupie Tracker"

// Configuraton globale
var appConfig *config.Config

// Création du templates
func CreateTemplates(app *config.Config) {
	appConfig = app
}

// Injection des templates
func renderTemplates(w http.ResponseWriter, tmplName string, td *models.MainData) {
	templateCache := appConfig.TemplateCache
	tmpl, ok := templateCache[tmplName+".page.tmpl"]

	// En cas d'erreur d'injection
	if !ok {
		err := pkg.Error(http.StatusInternalServerError)
		renderTemplates(w, "error", &models.MainData{AppInfos: models.App{PageTitle: err["code"] + " " + err["msg"]}, Error: models.ErrorData{Code: err["code"], Message: err["msg"]}})
		return
	}

	buffer := new(bytes.Buffer)
	tmpl.Execute(buffer, td)
	buffer.WriteTo(w)
}

// Création des caches
func CreateTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}
	pages, err := filepath.Glob("./templates/*.page.tmpl")

	if err != nil {
		return cache, err
	}

	// Parcourir toutes les pages
	for _, page := range pages {
		name := filepath.Base(page)
		tmpl := template.Must(template.ParseFiles(page))

		layouts, err := filepath.Glob("./templates/layouts/*.layout.tmpl")
		if err != nil {
			return cache, err
		}

		if len(layouts) > 0 {
			tmpl.ParseGlob("./templates/layouts/*.layout.tmpl")
		}

		cache[name] = tmpl
	}

	return cache, nil
}
