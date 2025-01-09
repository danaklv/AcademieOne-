package utils

import (
	utils "aw/utils/asciiArt"
	"html/template"
	"net/http"
)

type PageData struct {
	Result string
}

func MainPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet { 
		w.WriteHeader(405)
		return
	}
	if r.URL.Path == "/favicon.ico" {
		w.WriteHeader(200)
		return
	}
	if r.URL.Path != "/" {
		w.WriteHeader(404)
		tmpl, err := template.ParseFiles("web/404.html")
		if err != nil {
			utils.Status500(w)
			return
		}
		data := PageData{Result: "Not Found"} 
		if err := tmpl.Execute(w, data); err != nil {
			utils.Status500(w)
			return
		}
		return
	}

	tmpl, err := template.ParseFiles("web/index.html")
	if err != nil {
		utils.Status500(w)
		return
	}
	if err := tmpl.Execute(w, nil); err != nil {
		utils.Status500(w)
		return
	}
}
