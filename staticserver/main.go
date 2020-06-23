package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.StrictSlash(true)
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./public/"))))
	h := &http.Server{
		Addr:    ":3000",
		Handler: r,
	}
	h.ListenAndServe()
}
