package check

import (
	"gt/backend/models"
	"html/template"
	"net/http"
)

func Status500(w http.ResponseWriter) {
	w.WriteHeader(500)
	tmpl, err := template.ParseFiles("frontend/500.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data := models.PageData{Result: "Internal Server Error"}
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
func Status400(w http.ResponseWriter) {
	w.WriteHeader(400)
	tmpl, err := template.ParseFiles("frontend/400.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data := models.PageData{Result: "Bad Request"}
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func Status404(w http.ResponseWriter) {
	w.WriteHeader(404)
	tmpl, err := template.ParseFiles("frontend/404.html")
	if err != nil {
		Status500(w)
		return
	}
	data := models.PageData{Result: "Not Found"}
	if err := tmpl.Execute(w, data); err != nil {
		Status500(w)
		return
	}

}
