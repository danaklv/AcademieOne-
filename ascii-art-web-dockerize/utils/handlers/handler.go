package utils

import (
	"net/http"
)

func Handler() {
	http.HandleFunc("/", MainPageHandler)
	http.HandleFunc("/styles.css", CssHandler)
	http.HandleFunc("/ascii-art", AsciiArtHandler)
	http.HandleFunc("/download-sample", DownloadSampleHandler)
}
   