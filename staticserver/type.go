package main

import "log"

type Give struct {
	Types string
	Code  int
	Bytes []byte
}

type Data struct {
	Data  map[string]Give
	Types map[string]string
	Error []byte
}

func existe(e interface{}) {
	if e != nil {
	}
}

func NewData() (D *Data) {
	html, err := Asset("public/html/404.html")
	if err != nil {
		log.Fatal(err)
	}
	return &Data{
		Data: make(map[string]Give),
		Types: map[string]string{
			"css": "text/css",
			"png": "image/png",
		},
		Error: html,
	}
}
