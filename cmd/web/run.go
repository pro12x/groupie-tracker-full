package main

import (
	"bim/config"
	"bim/internal/handlers"
	"fmt"
	"net/http"
	"os"
)

func Run(tab []string) {
	if len(tab) == 0 {
		router := http.NewServeMux()
		var appConfig config.Config // Appel des éléments de configuration

		templateCache, err := handlers.CreateTemplateCache() // Création de cache pour switcher entre les pages

		if err != nil { // Retour en cas d'erreur
			panic(err)
		}

		appConfig.TemplateCache = templateCache // Création du cache de la vue
		appConfig.Port = ":1112"                // Configuration du port
		appConfig.StaticDir = "assets"          // Configuration des fichiers statiques
		handlers.CreateTemplates(&appConfig)    // Création de la vue

		// Passage des fichiers statiques
		statics := http.FileServer(http.Dir(appConfig.StaticDir))
		router.Handle("/static/", http.StripPrefix("/static/", statics))

		// Passages du gestionnaires des templates
		router.HandleFunc("/", handlers.Home)
		router.HandleFunc("/artists", handlers.Artists)
		router.HandleFunc("/artist/", handlers.Artist)
		router.HandleFunc("/list", handlers.List)

		// Lancement du serveur d'application
		fmt.Println("Server started on port http://localhost" + appConfig.Port)
		http.ListenAndServe(appConfig.Port, router)
		os.Exit(0) // Sortie du programme
	}
	fmt.Println("---<< Too many arguments, please try again >>---")
}
