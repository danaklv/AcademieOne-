package app

import (
	"aw/utils/handlers"
	"fmt"
	"net/http"
)

func Run() {

	utils.Handler()
	fmt.Println("http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
