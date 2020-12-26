package main

import (
	"log"
	"net/http"
)

type Give struct {
	Types string
	Code  int
	Bytes []byte
}

type Connect struct {
	connected bool
	user      string
	pwd       string
	jwtKey    []byte
	token     string
}

type H struct {
	H      func(w http.ResponseWriter, r *http.Request)
	Method []string
	Val  []int
}

type Oauth struct {

}

type Data struct {
	Url     map[string]H
	Data    map[string]Give
	Types   map[string]string
	Connect *Connect
	Oauth	*Oauth
	Error   []byte
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
		Url:  make(map[string]H),
		Types: map[string]string{
			"css": "text/css",
			"png": "image/png",
		},
		Connect: NewConnect(),
		Error:   html,
	}
}
