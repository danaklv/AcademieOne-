package utils

import (
	"aw/utils/asciiArt"
	"io"
	"net/http"
	"os"
	"strconv"
)

func DownloadSampleHandler(w http.ResponseWriter, r *http.Request) {
	file, err := os.OpenFile("sample.txt", os.O_RDWR, 0644)
	if err != nil {
		utils.Status500(w)
		return
	}
	defer file.Close()

	
	fileInfo, err := file.Stat()
	if err != nil {
		utils.Status500(w)
		return
	}

	w.Header().Set("Content-Disposition", "attachment; filename=sample.txt")
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Content-Length", strconv.FormatInt(fileInfo.Size(), 10))

	_, err = io.Copy(w, file)
	if err != nil {
		utils.Status500(w)
		return
	}
}
