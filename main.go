package main

import (
	"hangman-web/handlers"
	"log"
	"net/http"
)

func main() {
	// Servir les fichiers statiques
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	// Définir les routes
	http.HandleFunc("/promo", handlers.PromoHandler)
	http.HandleFunc("/change", handlers.ChangeHandler)
	http.HandleFunc("/user/form", handlers.UserFormHandler)
	http.HandleFunc("/user/treatment", handlers.UserTreatmentHandler)
	http.HandleFunc("/user/display", handlers.UserDisplayHandler)

	// Lancer le serveur
	log.Println("Serveur démarré sur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
