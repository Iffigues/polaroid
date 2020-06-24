package main

import (
	"net/http"
	//"strings"
)

func (a *Data)AddData(url string) (err error, val []byte) {
	println(url)
	return
}

func (a *Data)ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		w.Write([]byte("no file give"))
		return
	}
	if val, ok := a.Data[r.URL.Path]; ok {
		if val.Err != nil {
			w.Write([]byte(val.Err.Error()))
			return
		}
		w.Write(val.Bytes)
		return
	}
	err, data := a.AddData(r.URL.Path);
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	w.Write(data)
}

func main() {
	a := NewData()
	http.Handle("/", a)
	http.ListenAndServe(":3006", nil)
}
