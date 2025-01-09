package utils

import "net/http"

func CssHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web/styles.css")
}


