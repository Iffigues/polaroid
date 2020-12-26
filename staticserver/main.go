package main

import (
	"net/http"
	"strings"
)

func (a *Data) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	url := strings.Split(r.URL.Path, "/")
	if len(url) > 1 && url[1] == "public"  {
		if r.Method == "GET" {
			a.File(w, r)
			return
		}
	}
	for key, val := range a.Url {
		urls := strings.Split(key, "/")
		if len(url) == len(urls) {
			yes := true
			for i := 0; i < len(url); i = i + 1 {
				if urls[i] != "*" {
					if urls[i] != url[i] {
						yes = false
						break;
					}
				}
			}
			if yes {
				for _, valeur := range val.Method {
					if r.Method == valeur {
						val.H(w, r)
						return
					}
				}
				return
			}
		}
	}
	hello(w, r)
}

func main() {
	a := NewData()
	a.HandleFunc("/", []string{"GET"}, hello)
	a.HandleFunc("/upload", []string{"POST"}, Upload)
	a.HandleFunc("/public$", []string{"GET"}, a.File);
	http.Handle("/", a)
	http.ListenAndServe(":3006", nil)
}
