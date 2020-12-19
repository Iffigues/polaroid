package main

import (
	"errors"
	"net/http"
	"path/filepath"
	"strings"
)

func (a *Data) AddData(url string) (rr Give) {
	code := 200
	val, err := Asset(url[1:])
	ext := filepath.Ext(url[1:])
	types := "application/octet-stream"
	if len(ext) > 0 {
		ext = ext[1:]
	}
	if vals, ok := a.Types[ext]; ok {
		types = vals
	}
	if val == nil && err == nil {
		err = errors.New("empty data")
	}
	if err != nil {
		code = 404
		types = "text/html; charset=utf-8"
		val = a.Error
	}
	rr = Give{
		Code:  code,
		Types: types,
		Bytes: val,
	}
	a.Data[url] = rr
	return
}

func (a *Data) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	url := strings.Split(r.URL.Path, "/")
	for key, val := range a.Url {
		urls := strings.Split(key, "/")
		if len(url) == len(urls) {
			yes := true
			for i := 0; i < len(url); i = i + 1 {
				if urls[i] != "*" {
					if urls[i] != url[i] {
						yes = false
					}
				}
			}
			if yes {
				val.H(w, r)
				return
			}
		}
	}
	if val, ok := a.Url[r.URL.Path]; ok {
		val.H(w, r)
		return
	}
	if r.URL.Path == "/upload" {
		//		grap()
	}
	if strings.HasPrefix(r.URL.Path, "/public") {
		j := strings.Split(r.URL.Path, "/")
		if j[1] != "public" {
			hello(w, r)
			return
		}
		if val, ok := a.Data[r.URL.Path]; ok {
			w.Header().Set("Content-type", val.Types)
			w.WriteHeader(val.Code)
			w.Write(val.Bytes)
			return
		}
		val := a.AddData(r.URL.Path)
		w.Header().Set("Content-type", val.Types)
		w.WriteHeader(val.Code)
		w.Write(val.Bytes)
		return
	}
	hello(w, r)
}

func main() {
	a := NewData()
	a.HandleFunc("/", []string{"GET"}, hello)
	http.Handle("/", a)
	http.ListenAndServe(":3006", nil)
}
