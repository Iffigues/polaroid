package main

import ()

func (a *Connect) ConnectMe() (b bool) {
	return
}

func NewConnect() (a *Connect) {
	a = new(Connect)
	a.user = "css"
	a.pwd = "Mince1234"
	if !a.ConnectMe() {
		return
	}
	return
}
