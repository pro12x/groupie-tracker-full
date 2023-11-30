package config

import "text/template"

type Config struct {
	TemplateCache map[string]*template.Template // Cache
	Port          string                        // Port
	StaticDir     string                        // Répertoire des fichiers statiques
}
