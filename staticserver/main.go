package main

import (
	"net/http"
	"fmt"
	"errors"
	//"strings"
)

func (a *Data)AddData(aa string) (err error, val []byte) {

	var g Give
	e := len(aa)
	
	return
}

func (a *Data)ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if val, ok := a.Data[r.URL.Path]; ok {
		if val.Err != nil {
			fmt.Fprintf(w, val.Err.Error())
			return
		}
		fmt.Fprintf(w, val.Bytes)
		return
	}
	err, data := a.AddData(r.URL.Path);
	if err != nil {
	}
	fmt.Fprintf(w, data)
}

func main() {
	a := NewData()
	http.Handle("/", a)
	http.ListenAndServe(":3006", nil)
}
