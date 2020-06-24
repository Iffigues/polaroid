package main

import (
	"net/http"
	"fmt"
	"strings"
)

func hello(w http.ResponseWriter, r *http.Request) {
	e := r.URL.Path
	obj := e[1:]
	array := strings.Split(obj, "/")
	l := array[0]
	fmt.Fprintf(w, l)
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":3006", nil)
}
