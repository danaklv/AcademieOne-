package app

import (
	"fmt"
	"gt/backend/handler"
	"net/http"
)

func Run() {
	handler.Handler()
	fmt.Println("http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
