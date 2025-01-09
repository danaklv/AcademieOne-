package handler

import (
	"gt/backend/check"
	"gt/backend/models"
	"gt/backend/modify"
	"html/template"
	"net/http"
	"strconv"
)

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		check.Status400(w)
		return
	}
	intId, _ := strconv.Atoi(id)
	if intId < 1 || intId > 52 {
		check.Status404(w)
		return
	}
	artist, err := modify.GetArtistById(int64(intId), w)
	if err != nil {
		check.Status500(w)
		return
	}
	relations, err := modify.GetRelations(int64(intId), w)
	if err != nil {
		check.Status500(w)
		return
	}
	locations, err := modify.GetLocations(int64(intId), w)
	if err != nil {
		check.Status500(w)
		return
	}
	consertDates, err := modify.GetConcertDates(int64(intId), w)
	if err != nil {
		check.Status500(w)
		return
	}

	data := models.ArtistTemplate{
		Artist:       artist,
		Relations:    relations,
		Locations:    locations,
		ConcertDates: consertDates,
	}
	tmpl, err := template.ParseFiles("frontend/artist.html")
	if err != nil {
		check.Status500(w)
		return
	}

	if err := tmpl.Execute(w, data); err != nil {
		check.Status500(w)
		return
	}

}
