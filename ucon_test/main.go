package main

import (
	"net/http"

	"github.com/favclip/ucon"
)

func main() {
	ucon.Orthodox()

	ucon.HandleFunc("GET", "/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	ucon.ListenAndServe(":8080")
}
