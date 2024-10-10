package handlers

import (
	"html/template"
	"net/http"
	"regexp"
)

type User struct {
	Nom      string
	Prenom   string
	DateNais string
	Sexe     string
}

var user User

func UserFormHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/layout.html", "templates/user_form.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "layout", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func UserTreatmentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/user/form", http.StatusSeeOther)
		return
	}

	nom := r.FormValue("nom")
	prenom := r.FormValue("prenom")
	dateNais := r.FormValue("date_nais")
	sexe := r.FormValue("sexe")

	if !isValidName(nom) || !isValidName(prenom) || !isValidSexe(sexe) {
		http.Redirect(w, r, "/error", http.StatusSeeOther)
		return
	}

	user = User{
		Nom:      nom,
		Prenom:   prenom,
		DateNais: dateNais,
		Sexe:     sexe,
	}

	http.Redirect(w, r, "/user/display", http.StatusSeeOther)
}

func UserDisplayHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/layout.html", "templates/user_display.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "layout", user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func isValidName(name string) bool {
	matched, _ := regexp.MatchString("^[a-zA-Z]{1,32}$", name)
	return matched
}

func isValidSexe(sexe string) bool {
	return sexe == "masculin" || sexe == "féminin" || sexe == "autre"
}