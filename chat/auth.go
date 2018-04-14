package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

type authHandler struct {
	next http.Handler
}

func (h *authHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if _, err := r.Cookie("auth"); err == http.ErrNoCookie {
		w.Header().Set("Location", "/login")

		w.WriteHeader(http.StatusTemporaryRedirect)
	} else if err != nil {
		panic(err.Error())
	} else {
		h.next.ServeHTTP(w, r)
	}
}

func MustAuth(handler http.Handler) http.Handler {
	return &authHandler{next: handler}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	segs := strings.Split(r.URL.Path, "/")
	action := segs[2]
	if action == "" {
		w.WriteHeader(http.StatusNotFound)
		log.Fatal("action is empty")
	}

	provider := segs[3]
	if provider == "" {
		w.WriteHeader(http.StatusNotFound)
		log.Fatal("provider is empty")
	}

	switch action {
	case "login":
		log.Println("TODO: Login Process", provider)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "%s is unsupported", action)
	}
}
