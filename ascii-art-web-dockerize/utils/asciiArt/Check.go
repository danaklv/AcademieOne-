package utils

import (
	"aw/utils/models"
	"hash/fnv"
	"html/template"
	"net/http"
)

func CheckFile(style string) bool {
	hasher := fnv.New32a()
	hasher.Write([]byte(style))
	hashValue := hasher.Sum32()
	if hashValue != 1766917683 && hashValue != 4281396044 && hashValue != 3930937207 {
	
		return false
	}
	return true
}

func BadRequestHandler(input string) int {
	for _, r := range input {
		if (r < 32 || r > 127) && r != 10 && r != 13 {
			return 1
		}
	}
	return 0
}

func Status500(w http.ResponseWriter)  {
	w.WriteHeader(500)
		tmpl, err := template.ParseFiles("web/500.html")
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
	tmpl, err := template.ParseFiles("web/400.html")
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



