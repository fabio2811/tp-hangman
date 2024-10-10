package handlers

import (
	"html/template"
	"net/http"
)

type Student struct {
	Nom    string
	Prenom string
	Age    int
	Sexe   string
}

type Classe struct {
	Nom         string
	Filiere     string
	Niveau      string
	NbEtudiants int
	Etudiants   []Student
}

func PromoHandler(w http.ResponseWriter, r *http.Request) {
	classe := Classe{
		Nom:         "B1 Informatique",
		Filiere:     "Informatique",
		Niveau:      "Bachelor 1",
		NbEtudiants: 3,
		Etudiants: []Student{
			{Nom: "Dupont", Prenom: "Jean", Age: 20, Sexe: "masculin"},
			{Nom: "Martin", Prenom: "Marie", Age: 19, Sexe: "féminin"},
			{Nom: "Durand", Prenom: "Alex", Age: 21, Sexe: "autre"},
		},
	}

	tmpl, err := template.ParseFiles("templates/layout.html", "templates/promo.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "layout", classe)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}