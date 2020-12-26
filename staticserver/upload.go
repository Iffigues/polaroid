package main

import (
	"net/http"
	"strings"
	"encoding/base64"
)


func Upload(w http.ResponseWriter, r *http.Request) {
	i := r.Header.Get("Authorization")
	auth := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
	ee , err := base64.StdEncoding.DecodeString(auth[1])
	if err != nil {
		return
	}
	 strings.Split(string(ee), ":")

	w.Write([]byte(i+"zzz"))
}
