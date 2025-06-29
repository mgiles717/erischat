package main

import (
	"net/http"

	"github.com/mgiles717/erischat/internal/handlers"
)

func main() {
	http.HandleFunc("/", handlers.Home)
	http.ListenAndServe(":8080", nil)
}
