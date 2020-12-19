package main

import (
	"net/http"
)

func (a *Data) HandleFunc(aa string, m []string, b func(w http.ResponseWriter, r *http.Request)) {
	method := []string{}
	for _, val := range m {
		if val == "GET" {
			method = append(method, http.MethodGet)
		}
	}
	a.Url[aa] = H{
		H:      b,
		Method: m,
	}
}

func (a *Data) checkout(b string, w http.ResponseWriter, r *http.Request) (err error) {
	if val, ok := a.Url[b]; ok {
		method := r.Method
		for _, value := range val.Method {
			if method == value {
				val.H(w, r)
				return
			}
		}
	}
	return
}
