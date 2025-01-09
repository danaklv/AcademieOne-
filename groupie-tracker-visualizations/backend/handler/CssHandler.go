package handler

import "net/http"

func CssHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "frontend/style.css")
}

func SecondCssHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "frontend/style2.css")
}

func ErrCssHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "frontend/styleErr.css")
}

