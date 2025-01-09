package handler

import (
	"gt/backend/check"
	"gt/backend/models"
	"gt/backend/modify"
	"html/template"
	"net/http"
)

func MainPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(405)
		return
	}
	if r.URL.Path == "/favicon.ico" {
		w.WriteHeader(200)
		return
	}
	if r.URL.Path != "/" && r.URL.Path != "/index.html" {
		check.Status404(w)
		return
	}

	artists, err := modify.GetArtists(w)
	if err != nil {
		check.Status500(w)
		return
	}
	data := struct {
		Artists []models.Artist
	}{
		Artists: artists,
	}

	tmpl, err := template.ParseFiles("frontend/index.html")
	if err != nil {
		check.Status500(w)
		return
	}
	if err := tmpl.Execute(w, data); err != nil {
		check.Status500(w)
		return
	}
}
