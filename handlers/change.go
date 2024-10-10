package handlers

import (
	"html/template"
	"net/http"
	"sync"
)

var (
	viewCount int
	mutex     sync.Mutex
)

func ChangeHandler(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	viewCount++
	currentCount := viewCount
	mutex.Unlock()

	data := struct {
		Count  int
		IsPair bool
	}{
		Count:  currentCount,
		IsPair: currentCount%2 == 0,
	}

	tmpl, err := template.ParseFiles("templates/layout.html", "templates/change.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "layout", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
