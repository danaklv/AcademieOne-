package handler

import (
	"net/http"
)

func Handler() {
	http.HandleFunc("/", MainPageHandler)
	http.HandleFunc("/style.css", CssHandler)
	http.HandleFunc("/style2.css", SecondCssHandler)
	http.HandleFunc("/styleErr.css", ErrCssHandler)
	http.HandleFunc("/artist", ArtistHandler)
}
