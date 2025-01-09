package utils

import (
	"aw/utils/asciiArt"
	"html/template"
	"io/ioutil"
	"net/http"
)

func AsciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost && r.Method != http.MethodGet { //проверка метода
		w.WriteHeader(405)
		return
	}
	
	input := r.FormValue("input")
	banner := r.FormValue("choice")
	
	
	errr := utils.BadRequestHandler(input) 
	
	if errr == 0 {
		is500 :=  utils.AsciiArt(w, input, banner)
		if !is500 {
			return
		}
		result, err := ioutil.ReadFile("sample.txt")
		if err != nil {
			utils.Status500(w)
			return
		}
		asciiArt := string(result)
		tmpl, err := template.ParseFiles("web/result.html")
		if err != nil {
			utils.Status500(w)
			return
		}
		data := PageData{Result: asciiArt}
		if err := tmpl.Execute(w, data); err != nil {
			utils.Status500(w)
			return
		}
	} else { 
		utils.Status400(w) //BadRequest 
	}
}

